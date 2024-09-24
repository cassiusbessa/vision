package com.github.cassiusbessa.vision.domain.service.ports.output;

import com.github.cassiusbessa.vision.domain.core.entities.Profile;

import java.util.UUID;

public interface ProfileRepository {

    Profile findByProfileId(UUID accountId);

    Profile findByAccountId(UUID accountId);

    void save(Profile profile);

    void update(Profile profile);
}
