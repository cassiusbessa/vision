package com.github.cassiusbessa.vision.domain.service;

import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.core.entities.Profile;
import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.service.dtos.*;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceAlreadyExistsException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceNotFoundException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ValidationException;
import com.github.cassiusbessa.vision.domain.service.mappers.ProfileDataMapper;
import com.github.cassiusbessa.vision.domain.service.ports.input.ProfileService;
import com.github.cassiusbessa.vision.domain.service.ports.output.AccountRepository;
import com.github.cassiusbessa.vision.domain.service.ports.output.ProfileRepository;
import com.github.cassiusbessa.vision.domain.service.ports.output.TagRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@Slf4j
public class ProfileServiceImpl implements ProfileService {

    private final ProfileDataMapper profileDataMapper;
    private final ProfileRepository profileRepository;
    private final AccountRepository accountRepository;
    private final TagRepository tagRepository;

    @Autowired
    public ProfileServiceImpl(ProfileDataMapper profileDataMapper, ProfileRepository profileRepository, AccountRepository accountRepository, TagRepository tagRepository) {
        this.profileDataMapper = profileDataMapper;
        this.profileRepository = profileRepository;
        this.accountRepository = accountRepository;
        this.tagRepository = tagRepository;
    }

    @Override
    public ProfileCreatedResponse createProfile(ProfileCreateCommand command) {
        log.info("Creating profile for account: {}", command.getAccountId());

        if (command.getAccountId() == null) {
            log.error("Account ID is required");
            throw new ValidationException("Account ID is required");
        }

        Account account = accountRepository.findById(command.getAccountId());
        if (account == null) {
            log.error("Account does not exist: {}", command.getAccountId());
            throw new ResourceNotFoundException("Account does not exist: " + command.getAccountId());
        }

        Profile foundProfile = profileRepository.findByAccountId(command.getAccountId());
        if (foundProfile != null) {
            log.error("Profile already exists for account: {}", command.getAccountId());
            throw new ResourceAlreadyExistsException("Profile already exists for account: " + command.getAccountId());
        }

        List<Tag> tags;
        if (command.getTechnologies() != null && !command.getTechnologies().isEmpty()) {
            tags = tagRepository.findAllById(command.getTechnologies());
        } else {
            tags = List.of();
        }
        if (tags.size() != command.getTechnologies().size()) {
            log.error("Some tags do not exist: {}, only found: {}", command.getTechnologies(), tags);
            throw new ResourceNotFoundException("Some tags do not exist: " + command.getTechnologies());
        }

        Profile profile = profileDataMapper.profileCreateCommandToProfile(command, account, tags);
        profile.validate();
        if (!profile.getFailureMessages().isEmpty()) {
            log.error("Profile creation failed: {}", profile.getFailureMessagesAsString());
            throw new ValidationException(profile.getFailureMessagesAsString());
        }

        profileRepository.save(profile);
        log.info("Profile created successfully for account: {}", command.getAccountId());
        return new ProfileCreatedResponse("Profile created successfully");
    }

    @Override
    public ProfileUpdatedResponse updateProfile(ProfileUpdateCommand command) {
        return null;
    }

    @Override
    public LoadProfileResponse loadProfileById(LoadProfileByIdQuery query) {
        return null;
    }

    @Override
    public LoadProfileResponse loadProfileByAccountId(LoadProfileByAccountIdQuery query) {

        Profile profile = profileRepository.findByAccountId(query.getId());
        if (profile == null) {
            log.error("Profile not found for account: {}", query.getId());
            throw new ResourceNotFoundException("Profile not found for account: " + query.getId());
        }

        return new LoadProfileResponse(profile, "Loaded profile successfully");

    }
}
