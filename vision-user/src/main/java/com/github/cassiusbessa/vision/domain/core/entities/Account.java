package com.github.cassiusbessa.vision.domain.core.entities;

import com.github.cassiusbessa.vision.domain.core.valueobjects.*;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

public class Account extends AggregateRoot<AccountId> {

    private final Email email;
    private final EncryptedPassword password;
    private final AccountLevel accountLevel;
    private final List<String> failureMessages = new ArrayList<>();


    public Account(AccountId id, Email email, EncryptedPassword password, AccountLevel accountLevel) {
        super.setId(id);
        this.email = email;
        this.password = password;
        this.accountLevel = accountLevel;
    }

    public Email getEmail() {
        return email;
    }

    public EncryptedPassword getPassword() {
        return password;
    }

    public AccountLevel getAccountLevel() {
        return accountLevel;
    }

    public List<String> getFailureMessages() {
        return failureMessages;
    }

    public String getFailureMessagesAsString() {
        return String.join(", ", failureMessages);
    }

    public void validate() {
        if (email == null || email.getValue().isEmpty()) {
            failureMessages.add("Email is required");
        }
        if (password == null || password.getValue().isEmpty()) {
            failureMessages.add("Password is required");
        }
        if (password != null && password.getValue().length() < 6) {
            failureMessages.add("Password must have at least 6 characters");
        }
        if (accountLevel == null) {
            failureMessages.add("Account level is required");
        }
    }

    public static Builder builder() {
        return new Builder();
    }

    public static final class Builder {
        private AccountId id;
        private Email email;
        private EncryptedPassword password;
        private AccountLevel accountLevel;

        private Builder() {
        }

        public Builder withId(UUID id) {
            this.id = new AccountId(id);
            return this;
        }

        public Builder withEmail(Email email) {
            this.email = email;
            return this;
        }

        public Builder withPassword(EncryptedPassword password) {
            this.password = password;
            return this;
        }

        public Builder withAccountLevel(AccountLevel accountLevel) {
            this.accountLevel = accountLevel;
            return this;
        }

        public Account build() {
            return new Account(id, email, password, accountLevel);
        }
    }
}
