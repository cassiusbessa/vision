package com.github.cassiusbessa.vision.dataaccess;

import com.github.cassiusbessa.vision.dataaccess.entities.AccountDataBaseEntity;
import com.github.cassiusbessa.vision.dataaccess.mappers.AccountDataBaseMapper;
import com.github.cassiusbessa.vision.dataaccess.repositories.AccountJpaRepository;
import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.service.ports.output.AccountRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Repository;

import java.util.Optional;


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
    public Account findByEmail(String email) {
        Optional<AccountDataBaseEntity> accountDataBaseEntity = accountRepository.findByEmail(email);

        return accountDataBaseEntity.map(accountDataBaseMapper::dbEntityToAccount).orElse(null);
    }

    @Override
    public void save(Account account) {
        AccountDataBaseEntity accountDataBaseEntity = accountDataBaseMapper.accountToDbEntity(account);
        accountRepository.save(accountDataBaseEntity);

    }
}
