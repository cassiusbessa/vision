package com.github.cassiusbessa.vision.http;

import com.github.cassiusbessa.vision.domain.core.exceptions.DomainException;
import com.github.cassiusbessa.vision.domain.service.dtos.AccountCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.AccountCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.AuthCredentials;
import com.github.cassiusbessa.vision.domain.service.dtos.AuthResponse;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceAlredyExistsException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceNotFoundException;
import com.github.cassiusbessa.vision.domain.service.exceptions.UnauthorizedException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ValidationException;
import com.github.cassiusbessa.vision.domain.service.ports.input.AccountService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/account")
public class AccountController {

    private final AccountService accountService;

    @Autowired
    public AccountController(AccountService accountService) {
        this.accountService = accountService;
    }

    @PostMapping()
    public ResponseEntity<AccountCreatedResponse> createAccount(@RequestBody AccountCreateCommand command) {
        try {
            AccountCreatedResponse response = accountService.createAccount(command);
            return new ResponseEntity<>(response, HttpStatus.CREATED);
        } catch (ValidationException | DomainException e) {
            return new ResponseEntity<>(new AccountCreatedResponse(e.getMessage()), HttpStatus.BAD_REQUEST);
        } catch (ResourceAlredyExistsException e) {
            return new ResponseEntity<>(new AccountCreatedResponse(e.getMessage()), HttpStatus.CONFLICT);
        } catch (Exception e) {
            return new ResponseEntity<>(new AccountCreatedResponse(e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PostMapping("/auth")
    public ResponseEntity<AuthResponse> authenticate(@RequestBody AuthCredentials credentials) {
        try {
            AuthResponse response = accountService.auth(credentials);
            return new ResponseEntity<>(response, HttpStatus.OK);
        } catch (ResourceNotFoundException e) {
            return new ResponseEntity<>(new AuthResponse(null, e.getMessage()), HttpStatus.NOT_FOUND);
        } catch (UnauthorizedException e) {
            return new ResponseEntity<>(new AuthResponse(null, e.getMessage()), HttpStatus.UNAUTHORIZED);
        } catch (Exception e) {
            return new ResponseEntity<>(new AuthResponse(null, e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}
