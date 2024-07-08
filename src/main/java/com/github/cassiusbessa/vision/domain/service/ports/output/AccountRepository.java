package com.github.cassiusbessa.vision.domain.service.ports.output;

import com.github.cassiusbessa.vision.domain.core.entities.Account;

import java.util.UUID;

public interface AccountRepository {

    Account findByEmail(String email);

    Account findById(UUID id);

    boolean existsById(UUID id);

    void save(Account account);

}
