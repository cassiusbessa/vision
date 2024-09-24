package com.github.cassiusbessa.vision.domain.service.ports.output.repositories;

import com.github.cassiusbessa.vision.domain.core.entities.Account;

import java.util.UUID;

public interface AccountRepository {

    Account findById(UUID id);

    boolean existsById(UUID id);

}
