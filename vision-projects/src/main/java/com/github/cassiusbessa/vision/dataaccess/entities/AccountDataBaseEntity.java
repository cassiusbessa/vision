package com.github.cassiusbessa.vision.dataaccess.entities;

import com.github.cassiusbessa.vision.domain.core.valueobjects.AccountLevel;
import jakarta.persistence.*;

import java.util.UUID;

@Entity
@Table(name = "accounts")
public class AccountDataBaseEntity {

    @Id
    @Column(name = "id", nullable = false, unique = true)
    private UUID id;

    @Column(name = "email", nullable = false, unique = true)
    private String email;

    @Column(name = "password", nullable = false)
    private String password;

    @Enumerated(EnumType.STRING)
    @Column(name = "account_level", nullable = false)
    private AccountLevel accountLevel;

    public AccountDataBaseEntity() {
    }

    public AccountDataBaseEntity(UUID id, String email, String password) {
        this.id = id;
        this.email = email;
        this.password = password;
        accountLevel = AccountLevel.FREE;
    }

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword() { return password;  }

    public void setPassword(String password) {
        this.password = password;
    }

    public AccountLevel getAccountLevel() { return accountLevel; }

    public void setAccountLevel(AccountLevel accountLevel) { this.accountLevel = accountLevel; }
}
