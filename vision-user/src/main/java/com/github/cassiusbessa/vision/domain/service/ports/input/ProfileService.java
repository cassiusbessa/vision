package com.github.cassiusbessa.vision.domain.service.ports.input;

import com.github.cassiusbessa.vision.domain.service.dtos.profile.*;

public interface ProfileService {

    ProfileCreatedResponse createProfile(ProfileCreateCommand command);

    ProfileUpdatedResponse updateProfile(ProfileUpdateCommand command);

    LoadProfileResponse loadProfileById(LoadProfileByIdQuery query);

    LoadProfileResponse loadProfileByAccountId(LoadProfileByAccountIdQuery query);
}
