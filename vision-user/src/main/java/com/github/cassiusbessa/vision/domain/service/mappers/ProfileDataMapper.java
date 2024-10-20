package com.github.cassiusbessa.vision.domain.service.mappers;

import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.core.entities.Profile;
import com.github.cassiusbessa.vision.domain.core.entities.Project;
import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.service.dtos.profile.ProfileCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.profile.ProfileDTO;
import com.github.cassiusbessa.vision.domain.service.dtos.profile.ProfileUpdateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.tag.TagDTO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.HashSet;
import java.util.List;
import java.util.UUID;

@Component
public class ProfileDataMapper {

    private final ProjectMapper projectMapper;

    @Autowired
    public ProfileDataMapper(ProjectMapper projectMapper) {
        this.projectMapper = projectMapper;
    }

    public Profile profileCreateCommandToProfile(ProfileCreateCommand command, Account account, List<Tag> tags) {
        return Profile.builder()
                .withId(UUID.randomUUID())
                .withName(command.getName())
                .withTitle(command.getTitle())
								.withImage(command.getImage())
                .withDescription(command.getDescription())
                .withTechnologies(new HashSet<>(tags))
                .withAccount(account)
                .withLink(command.getLink())
                .build();
    }

    public ProfileDTO profileToProfileDTO(Profile profile) {

        var profileDTOTags = new HashSet<TagDTO>();

        for (var tag : profile.getTechnologies()) {
            profileDTOTags.add(new TagDTO(tag.getId().getValue(), tag.getName()));
        }

        return new ProfileDTO(
                profile.getId().getValue(),
                profile.getName(),
                profile.getImage(),
                profile.getTitle(),
                profile.getDescription(),
                profileDTOTags,
                projectMapper.projectToProjectDTO(profile.getStarProject()),
                profile.getLink()
        );
    }

    public Profile profileUpdateCommandToProfile(ProfileUpdateCommand command, Account account, List<Tag> tags, Project startProject) {
        return Profile.builder()
                .withId(command.getProfileId())
                .withName(command.getName())
                .withTitle(command.getTitle())
                .withDescription(command.getDescription())
                .withTechnologies(new HashSet<>(tags))
                .withStarProject(startProject)
                .withAccount(account)
                .withLink(command.getLink())
                .build();
    }
}
