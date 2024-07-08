package com.github.cassiusbessa.vision.domain.service.mappers;

import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.core.valueobjects.AccountLevel;
import com.github.cassiusbessa.vision.domain.core.valueobjects.Email;
import com.github.cassiusbessa.vision.domain.core.valueobjects.EncryptedPassword;
import com.github.cassiusbessa.vision.domain.core.valueobjects.Password;
import com.github.cassiusbessa.vision.domain.service.crypto.CryptoService;
import com.github.cassiusbessa.vision.domain.service.dtos.AccountCreateCommand;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.UUID;

@Component
public class AccountDataMapper {

    private final CryptoService cryptoService;

    @Autowired
    public AccountDataMapper(CryptoService cryptoService) {
        this.cryptoService = cryptoService;
    }

    public Account createAccountCommandToAccount(AccountCreateCommand command) {
        return Account.builder()
                .withId(UUID.randomUUID())
                .withEmail(new Email(command.getEmail()))
                .withPassword(new EncryptedPassword(cryptoService.encrypt(command.getPassword())))
                .withAccountLevel(AccountLevel.FREE)
                .build();
    }
}
