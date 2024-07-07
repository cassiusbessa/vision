package com.github.cassiusbessa.vision.domain.service;

import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.service.crypto.CryptoService;
import com.github.cassiusbessa.vision.domain.service.dtos.AccountCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.AccountCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.AuthCredentials;
import com.github.cassiusbessa.vision.domain.service.dtos.AuthResponse;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceAlredyExistsException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceNotFoundException;
import com.github.cassiusbessa.vision.domain.service.exceptions.UnauthorizedException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ValidationException;
import com.github.cassiusbessa.vision.domain.service.mappers.AccountDataMapper;
import com.github.cassiusbessa.vision.domain.service.ports.input.AccountService;
import com.github.cassiusbessa.vision.domain.service.ports.output.AccountRepository;
import com.github.cassiusbessa.vision.domain.service.token.JwtTokenService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;



@Slf4j
@Service
public class AccountServiceImpl implements AccountService {

    private final AccountRepository accountRepository;

    private final AccountDataMapper accountDataMapper;

    private final CryptoService cryptoService;

    private final JwtTokenService tokenService;



    @Autowired
    public AccountServiceImpl(AccountRepository accountRepository, AccountDataMapper accountDataMapper, CryptoService cryptoService, JwtTokenService tokenService) {
        this.accountRepository = accountRepository;
        this.accountDataMapper = accountDataMapper;
        this.cryptoService = cryptoService;
        this.tokenService = tokenService;
    }

    @Override
    public AccountCreatedResponse createAccount(AccountCreateCommand command) {
        log.info("Creating account with email: {}", command.getEmail());

        Account foundAccount = accountRepository.findByEmail(command.getEmail());
        if (foundAccount != null) {
            log.error("Account already exists with email: {}", command.getEmail());
            throw new ResourceAlredyExistsException("Account already exists with email: " + command.getEmail());
        }

        Account account = accountDataMapper.createAccountCommandToAccount(command);

        account.validate();
        if (!account.getFailureMessages().isEmpty()) {
            log.error("Account creation failed: {}", account.getFailureMessagesAsString());
            throw new ValidationException(account.getFailureMessagesAsString());
        }

        accountRepository.save(account);
        log.info("Account created successfully");
        return new AccountCreatedResponse("Account created successfully");
    }

    @Override
    public AuthResponse auth(AuthCredentials credentials) {
        log.info("Logging in with email: {}", credentials.getEmail());

        Account account = accountRepository.findByEmail(credentials.getEmail());
        if (account == null) {
            log.error("Account not found with email: {}", credentials.getEmail());
            throw new ResourceNotFoundException("Account not found with email: " + credentials.getEmail());
        }
        if (!cryptoService.matches(credentials.getPassword(), account.getPassword().getValue())) {
            log.error("Invalid password for account with email: {}", credentials.getEmail());
            throw new UnauthorizedException("Invalid credentials");
        }
        log.info("Logged in successfully with email: {}", credentials.getEmail());
        return new AuthResponse(tokenService.generateToken(account.getId().getValue()), "Logged in successfully");
    }
}
