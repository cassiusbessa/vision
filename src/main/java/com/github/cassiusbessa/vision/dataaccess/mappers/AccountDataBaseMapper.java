package com.github.cassiusbessa.vision.dataaccess.mappers;

import com.github.cassiusbessa.vision.dataaccess.entities.AccountDataBaseEntity;
import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.core.valueobjects.Email;
import com.github.cassiusbessa.vision.domain.core.valueobjects.EncryptedPassword;
import com.github.cassiusbessa.vision.domain.core.valueobjects.Password;
import org.springframework.stereotype.Component;

@Component
public class AccountDataBaseMapper {

    public Account dbEntityToAccount(AccountDataBaseEntity dbEntity) {
        return Account.builder()
                .withId(dbEntity.getId())
                .withEmail(new Email(dbEntity.getEmail()))
                .withPassword(new EncryptedPassword(dbEntity.getPassword()))
                .build();
    }

    public AccountDataBaseEntity accountToDbEntity(Account account) {
        return new AccountDataBaseEntity(account.getId().getValue(), account.getEmail().getValue(), account.getPassword().getValue());
    }
}
