package com.github.cassiusbessa.vision.domain.service;

import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.core.entities.Profile;
import com.github.cassiusbessa.vision.domain.core.entities.Project;
import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.core.valueobjects.AccountId;
import com.github.cassiusbessa.vision.domain.service.dtos.profile.*;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceAlreadyExistsException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceNotFoundException;
import com.github.cassiusbessa.vision.domain.service.exceptions.UnauthorizedException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ValidationException;
import com.github.cassiusbessa.vision.domain.service.mappers.ProfileDataMapper;
import com.github.cassiusbessa.vision.domain.service.ports.input.ProfileService;
import com.github.cassiusbessa.vision.domain.service.ports.output.AccountRepository;
import com.github.cassiusbessa.vision.domain.service.ports.output.ProfileRepository;
import com.github.cassiusbessa.vision.domain.service.ports.output.ProjectRepository;
import com.github.cassiusbessa.vision.domain.service.ports.output.TagRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.UUID;

@Service
@Slf4j
public class ProfileServiceImpl implements ProfileService {

    private final ProfileDataMapper profileDataMapper;
    private final ProfileRepository profileRepository;
    private final AccountRepository accountRepository;
    private final TagRepository tagRepository;
    private final ProjectRepository projectRepository;

    @Autowired
    public ProfileServiceImpl(ProfileDataMapper profileDataMapper, ProfileRepository profileRepository, AccountRepository accountRepository, TagRepository tagRepository, ProjectRepository projectRepository) {
        this.profileDataMapper = profileDataMapper;
        this.profileRepository = profileRepository;
        this.accountRepository = accountRepository;
        this.tagRepository = tagRepository;
        this.projectRepository = projectRepository;
    }

    @Override
    public ProfileCreatedResponse createProfile(ProfileCreateCommand command) {
        log.info("Creating profile for account: {}", command.getAccountId());

        Account account = getAccount(command.getAccountId());

        Profile foundProfile = profileRepository.findByAccountId(command.getAccountId());
        if (foundProfile != null) {
            log.error("Profile already exists for account: {}", command.getAccountId());
            throw new ResourceAlreadyExistsException("Profile already exists for account: " + command.getAccountId());
        }

				checkProfileLinkExists(command.getLink(), null);

        List<Tag> tags = getTags(command.getTechnologies());

        Profile profile = profileDataMapper.profileCreateCommandToProfile(command, account, tags);
        validateProfile(profile);

        profileRepository.save(profile);
        log.info("Profile created successfully for account: {}", command.getAccountId());
        return new ProfileCreatedResponse("Profile created successfully");
    }

        @Override
        public ProfileUpdatedResponse updateProfile(ProfileUpdateCommand command) {
            log.info("Updating profile: {}", command.getProfileId());

            Account account = getAccount(command.getAccountId());

						checkProfileLinkExists(command.getLink(), account.getId());
						
            List<Tag> tags = getTags(command.getTechnologies());

            Project startProject = projectRepository.findByProjectId(command.getStarProjectId());

            if (startProject == null) {
                log.error("Project does not exist: {}", command.getStarProjectId());
                throw new ResourceNotFoundException("Project does not exist: " + command.getStarProjectId());
            }
            if (!startProject.getAccount().getId().equals(account.getId())) {
                log.error("Project does not belong to account: {}", command.getStarProjectId());
                throw new UnauthorizedException("Project does not belong to account: " + command.getStarProjectId());
            }

            Profile profile = profileDataMapper.profileUpdateCommandToProfile(command, account, tags, startProject);
            validateProfile(profile);

            profileRepository.update(profile);
            log.info("Profile updated successfully: {}", command.getProfileId());
            return new ProfileUpdatedResponse("Profile updated successfully");
        }

    @Override
    public LoadProfileResponse loadProfileById(LoadProfileByIdQuery query) {
        log.info("Loading profile: {}", query.getId());

        System.out.println("Loading profile: " + query.getId());

        Profile profile = profileRepository.findByProfileId(query.getId());
        if (profile == null) {
            log.error("Profile not found: {}", query.getId());
            throw new ResourceNotFoundException("Profile not found: " + query.getId());
        }

        ProfileDTO profileDTO = profileDataMapper.profileToProfileDTO(profile);
        return new LoadProfileResponse(profileDTO, "Loaded profile successfully");
    }

    @Override
    public LoadProfileResponse loadProfileByAccountId(LoadProfileByAccountIdQuery query) {
        log.info("Loading profile for account: {}", query.getId());

        Profile profile = profileRepository.findByAccountId(query.getId());
        if (profile == null) {
            log.error("Profile not found for account: {}", query.getId());
            throw new ResourceNotFoundException("Profile not found for account: " + query.getId());
        }

        ProfileDTO profileDTO = profileDataMapper.profileToProfileDTO(profile);
        return new LoadProfileResponse(profileDTO, "Loaded profile successfully");
    }

		@Override
		public LoadProfileResponse loadProfileByLink(LoadProfileByLinkQuery query) {
				log.info("Loading profile by link: {}", query.link());

				Profile profile = profileRepository.findByLink(query.link());
				if (profile == null) {
						log.error("Profile not found by link: {}", query.link());
						throw new ResourceNotFoundException("Profile not found by link: " + query.link());
				}

				ProfileDTO profileDTO = profileDataMapper.profileToProfileDTO(profile);
				return new LoadProfileResponse(profileDTO, "Loaded profile successfully");
		}


    private Account getAccount(UUID accountId) {
        if (accountId == null) {
            log.error("Account ID is required");
            throw new ValidationException("Account ID is required");
        }
        Account account = accountRepository.findById(accountId);
        if (account == null) {
            log.error("Account does not exist: {}", accountId);
            throw new ResourceNotFoundException("Account does not exist: " + accountId);
        }
        return account;
    }

    private void validateProfile(Profile profile) {
        profile.validate();
        if (!profile.getFailureMessages().isEmpty()) {
            log.error("Profile update failed: {}", profile.getFailureMessagesAsString());
            throw new ValidationException(profile.getFailureMessagesAsString());
        }
    }

		private void checkProfileLinkExists(String link, AccountId accountId) {
			if (accountId == null) {
				return;
			}
			Profile profile = profileRepository.findByLink(link);
			if (profile != null && !profile.getAccount().getId().equals(accountId)) {
				log.error("Profile link already exists: {}", link);
				throw new ResourceAlreadyExistsException("Profile link already exists: " + link);
			}
		}

    private List<Tag> getTags(List<UUID> tagIds) {
        List<Tag> tags = tagRepository.findAllById(tagIds);
        validateTags(tags, tagIds);
        return tags;
    }

    private void validateTags(List<Tag> tags, List<UUID> tagIds) {
        if (tags.size() != tagIds.size()) {
            log.error("Some tags do not exist: {}, only found: {}", tagIds, tags);
            throw new ResourceNotFoundException("Some tags do not exist: " + tagIds);
        }
    }
}
