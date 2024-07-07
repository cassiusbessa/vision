package com.github.cassiusbessa.vision.domain.service.ports.input;

import com.github.cassiusbessa.vision.domain.service.dtos.AccountCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.AccountCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.LoginCredentials;
import com.github.cassiusbessa.vision.domain.service.dtos.LoginResponse;

public interface AccountService {

    AccountCreatedResponse createAccount(AccountCreateCommand command);

    LoginResponse login(LoginCredentials credentials);
}
