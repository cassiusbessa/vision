package com.github.cassiusbessa.vision.http;

import com.github.cassiusbessa.vision.domain.core.exceptions.DomainException;
import com.github.cassiusbessa.vision.domain.service.dtos.account.AccountCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.account.AccountCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.account.LoadedAccountResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.auth.AuthCredentials;
import com.github.cassiusbessa.vision.domain.service.dtos.auth.AuthResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.auth.MeResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.profile.LoadProfileByAccountIdQuery;
import com.github.cassiusbessa.vision.domain.service.dtos.profile.LoadProfileResponse;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceAlreadyExistsException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceNotFoundException;
import com.github.cassiusbessa.vision.domain.service.exceptions.UnauthorizedException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ValidationException;
import com.github.cassiusbessa.vision.domain.service.ports.input.AccountService;
import com.github.cassiusbessa.vision.domain.service.ports.input.ProfileService;
import com.github.cassiusbessa.vision.domain.service.ports.input.TokenService;


import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@RestController
@RequestMapping("/account")
public class AccountController {

    private final AccountService accountService;
		private final ProfileService profileService;
		private final TokenService tokenService;

    @Autowired
    public AccountController(AccountService accountService, TokenService tokenService, ProfileService profileService) {
			this.accountService = accountService;
			this.profileService = profileService;
			this.tokenService = tokenService;
    }

    @PostMapping()
    public ResponseEntity<AccountCreatedResponse> createAccount(@RequestBody AccountCreateCommand command) {
			try {
					AccountCreatedResponse response = accountService.createAccount(command);
					return new ResponseEntity<>(response, HttpStatus.CREATED);
			} catch (ValidationException | DomainException e) {
					return new ResponseEntity<>(new AccountCreatedResponse(e.getMessage()), HttpStatus.BAD_REQUEST);
			} catch (ResourceAlreadyExistsException e) {
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

		@GetMapping("/me")
		public ResponseEntity<MeResponse> loadMe(@RequestHeader("Authorization") String token) throws UnauthorizedException {
			try {
				UUID accountId = tokenService.getAccountId(token);
				LoadedAccountResponse loadedAccount = accountService.loadAccountById(accountId);
				LoadProfileByAccountIdQuery query = new LoadProfileByAccountIdQuery(accountId);
				LoadProfileResponse profile = profileService.loadProfileByAccountId(query);
				MeResponse response = new MeResponse(loadedAccount.account(), profile.getProfile(), "Loaded account successfully");
				return ResponseEntity.ok(response);
			} catch (ResourceNotFoundException e) {
				return new ResponseEntity<>(new MeResponse(null, null, e.getMessage()), HttpStatus.NOT_FOUND);
			} catch (UnauthorizedException e) {
				return new ResponseEntity<>(new MeResponse(null, null, e.getMessage()), HttpStatus.UNAUTHORIZED);
			} catch (Exception e) {
				return new ResponseEntity<>(new MeResponse(null, null, e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
			}
	}
}
