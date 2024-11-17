package com.github.cassiusbessa.vision.domain.service.ports.input;

import java.util.UUID;

import com.github.cassiusbessa.vision.domain.service.dtos.account.AccountCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.account.AccountCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.account.LoadedAccountResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.auth.AuthCredentials;
import com.github.cassiusbessa.vision.domain.service.dtos.auth.AuthResponse;

public interface AccountService {

    AccountCreatedResponse createAccount(AccountCreateCommand command);

    AuthResponse auth(AuthCredentials credentials);

		LoadedAccountResponse loadAccountById(UUID accountId);
}
