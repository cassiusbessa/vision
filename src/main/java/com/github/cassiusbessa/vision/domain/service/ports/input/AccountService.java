package com.github.cassiusbessa.vision.domain.service.ports.input;

import com.github.cassiusbessa.vision.domain.service.dtos.AccountCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.AccountCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.AuthCredentials;
import com.github.cassiusbessa.vision.domain.service.dtos.AuthResponse;

public interface AccountService {

    AccountCreatedResponse createAccount(AccountCreateCommand command);

    AuthResponse auth(AuthCredentials credentials);
}
