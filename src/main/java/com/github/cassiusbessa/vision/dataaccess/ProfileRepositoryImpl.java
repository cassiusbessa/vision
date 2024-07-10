package com.github.cassiusbessa.vision.dataaccess;

import com.github.cassiusbessa.vision.dataaccess.entities.ProfileDataBaseEntity;
import com.github.cassiusbessa.vision.dataaccess.mappers.ProfileDataBaseMapper;
import com.github.cassiusbessa.vision.dataaccess.repositories.ProfileJpaRepository;
import com.github.cassiusbessa.vision.domain.core.entities.Profile;
import com.github.cassiusbessa.vision.domain.service.ports.output.ProfileRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public class ProfileRepositoryImpl implements ProfileRepository {

    private final ProfileDataBaseMapper profileDataBaseMapper;
    private final ProfileJpaRepository profileRepository;

    @Autowired
    public ProfileRepositoryImpl(ProfileDataBaseMapper profileDataBaseMapper, ProfileJpaRepository profileRepository) {
        this.profileDataBaseMapper = profileDataBaseMapper;
        this.profileRepository = profileRepository;
    }

    @Override
    public Profile findByProfileId(UUID profileId) {
        ProfileDataBaseEntity profileDataBaseEntity = profileRepository.findById(profileId).orElse(null);
        if (profileDataBaseEntity == null) {
            return null;
        }
        return profileDataBaseMapper.dbEntityToProfile(profileDataBaseEntity);
    }

    @Override
    public Profile findByAccountId(UUID accountId) {
        ProfileDataBaseEntity profileDataBaseEntity = profileRepository.findByAccountId(accountId).orElse(null);
        if (profileDataBaseEntity == null) {
            return null;
        }
        return profileDataBaseMapper.dbEntityToProfile(profileDataBaseEntity);
    }

    @Override
    public void save(Profile profile) {
        profileRepository.save(profileDataBaseMapper.profileToDbEntity(profile));
    }

    @Override
    public void update(Profile profile) {

        ProfileDataBaseEntity profileDataBaseEntity = profileRepository.findByAccountId(profile.getAccount().getId().getValue()).orElse(null);
        if (profileDataBaseEntity == null) {
            return;
        }
        profileDataBaseMapper.updateProfile(profile, profileDataBaseEntity);
        profileRepository.save(profileDataBaseEntity);

    }
}
