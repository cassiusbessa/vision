package com.github.cassiusbessa.vision.domain.service.ports.output;

import com.github.cassiusbessa.vision.domain.core.entities.Account;

public interface AccountRepository {

    Account findByEmail(String email);

    void save(Account account);

}
