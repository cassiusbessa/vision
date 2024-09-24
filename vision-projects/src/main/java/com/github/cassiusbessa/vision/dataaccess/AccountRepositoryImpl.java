package com.github.cassiusbessa.vision.dataaccess;

import com.github.cassiusbessa.vision.dataaccess.entities.AccountDataBaseEntity;
import com.github.cassiusbessa.vision.dataaccess.mappers.AccountDataBaseMapper;
import com.github.cassiusbessa.vision.dataaccess.repositories.AccountJpaRepository;
import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.service.ports.output.repositories.AccountRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.Optional;
import java.util.UUID;


@Repository
public class AccountRepositoryImpl implements AccountRepository {

    private final AccountJpaRepository accountRepository;
    private final AccountDataBaseMapper accountDataBaseMapper;

    @Autowired
    public AccountRepositoryImpl(AccountJpaRepository accountRepository, AccountDataBaseMapper accountDataBaseMapper) {
        this.accountRepository = accountRepository;
        this.accountDataBaseMapper = accountDataBaseMapper;
    }

    @Override
    public Account findById(UUID id) {
        Optional<AccountDataBaseEntity> accountDataBaseEntity = accountRepository.findById(id);

        return accountDataBaseEntity.map(accountDataBaseMapper::dbEntityToAccount).orElse(null);
    }

    @Override
    public boolean existsById(UUID id) {
        return accountRepository.existsById(id);
    }

}
