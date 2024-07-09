package com.github.cassiusbessa.vision.dataaccess.mappers;

import com.github.cassiusbessa.vision.dataaccess.entities.ProfileDataBaseEntity;
import com.github.cassiusbessa.vision.domain.core.entities.Profile;
import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.core.valueobjects.ProfileId;
import com.github.cassiusbessa.vision.domain.core.valueobjects.TagId;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.stream.Collectors;

@Component
public class ProfileDataBaseMapper {

    private final TagDataBaseMapper tagDataBaseMapper;
    private final AccountDataBaseMapper accountMapper;
    private final ProjectDataBaseMapper projectMapper;

    @Autowired
    public ProfileDataBaseMapper(TagDataBaseMapper tagDataBaseMapper, AccountDataBaseMapper accountMapper, ProjectDataBaseMapper projectMapper) {
        this.tagDataBaseMapper = tagDataBaseMapper;
        this.accountMapper = accountMapper;
        this.projectMapper = projectMapper;
    }

    public Profile dbEntityToProfile(ProfileDataBaseEntity dbEntity) {

        return Profile.builder()
                .withId(dbEntity.getId())
                .withName(dbEntity.getName())
                .withTitle(dbEntity.getTitle())
                .withDescription(dbEntity.getDescription())
                .withTechnologies(
                        dbEntity.getTechnologies().stream()
                                .map(tag -> new Tag(new TagId(tag.getId()), tag.getName()))
                                .collect(Collectors.toSet())
                )
                .withImage(dbEntity.getImage())
                .withAccount(accountMapper.dbEntityToAccount(dbEntity.getAccount()))
                .build();
    }

    public ProfileDataBaseEntity profileToDbEntity(Profile profile) {
        return new ProfileDataBaseEntity(
                profile.getId().getValue(),
                profile.getName(),
                profile.getImage(),
                profile.getTitle(),
                profile.getDescription(),
                accountMapper.accountToDbEntity(profile.getAccount()),
                projectMapper.projectToDbEntity(profile.getStarProject()),
                profile.getTechnologies().stream().map(tagDataBaseMapper::tagToDbEntity).collect(Collectors.toSet())
        );
    }
}
