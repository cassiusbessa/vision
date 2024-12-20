package com.github.cassiusbessa.vision.http;

import com.github.cassiusbessa.vision.domain.core.exceptions.DomainException;
import com.github.cassiusbessa.vision.domain.service.dtos.profile.*;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceAlreadyExistsException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceNotFoundException;
import com.github.cassiusbessa.vision.domain.service.exceptions.UnauthorizedException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ValidationException;
import com.github.cassiusbessa.vision.domain.service.ports.input.ProfileService;
import com.github.cassiusbessa.vision.domain.service.ports.input.TokenService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.UUID;

@Slf4j
@RestController
@RequestMapping("/profile")
public class ProfileController {

    private final ProfileService profileService;
    private final TokenService tokenService;

    @Autowired
    public ProfileController(ProfileService profileService, TokenService tokenService) {
        this.profileService = profileService;
        this.tokenService = tokenService;
    }

    @PostMapping()
    public ResponseEntity<ProfileCreatedResponse> createProfile(@RequestBody ProfileCreateCommand command, @RequestHeader("Authorization") String token){
        try {
						UUID accountId = tokenService.getAccountId(token);
						command.setAccountId(accountId);
            ProfileCreatedResponse response = profileService.createProfile(command);
            return ResponseEntity.ok(response);
        } catch (ValidationException | DomainException e) {
            return new ResponseEntity<>(new ProfileCreatedResponse(e.getMessage()), HttpStatus.BAD_REQUEST);
        } catch (ResourceAlreadyExistsException e) {
            return new ResponseEntity<>(new ProfileCreatedResponse(e.getMessage()), HttpStatus.CONFLICT);
        }  catch (UnauthorizedException e) {
            return new ResponseEntity<>(new ProfileCreatedResponse(e.getMessage()), HttpStatus.UNAUTHORIZED);
        } catch (Exception e) {
            log.error("Error creating profile", e);
            return new ResponseEntity<>(new ProfileCreatedResponse(e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PutMapping()
    public ResponseEntity<ProfileUpdatedResponse> updateProfile(@RequestBody ProfileUpdateCommand command, @RequestHeader("Authorization") String token){
        try {
            if (!tokenService.getAccountId(token).equals(command.getAccountId())) {
                log.error("Unauthorized profile update, invalid token");
                return new ResponseEntity<>(new ProfileUpdatedResponse("Unauthorized"), HttpStatus.UNAUTHORIZED);
            }
            ProfileUpdatedResponse response = profileService.updateProfile(command);
            return ResponseEntity.ok(response);
        } catch (ResourceNotFoundException e) {
            return new ResponseEntity<>(new ProfileUpdatedResponse(e.getMessage()), HttpStatus.NOT_FOUND);
        } catch (ValidationException | DomainException e) {
            return new ResponseEntity<>(new ProfileUpdatedResponse(e.getMessage()), HttpStatus.BAD_REQUEST);
        } catch (ResourceAlreadyExistsException e) {
            return new ResponseEntity<>(new ProfileUpdatedResponse(e.getMessage()), HttpStatus.CONFLICT);
        }  catch (UnauthorizedException e) {
            return new ResponseEntity<>(new ProfileUpdatedResponse(e.getMessage()), HttpStatus.UNAUTHORIZED);
        } catch (Exception e) {
            log.error("Error updating profile", e);
            return new ResponseEntity<>(new ProfileUpdatedResponse(e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @GetMapping("/{id}")
    public ResponseEntity<LoadProfileResponse> loadProfileById(@PathVariable("id") String id) {
        try {
            LoadProfileResponse response = profileService.loadProfileById(new LoadProfileByIdQuery(
                    UUID.fromString(id)
            ));
            return ResponseEntity.ok(response);
        } catch (ResourceNotFoundException e) {
            return new ResponseEntity<>(new LoadProfileResponse(null, e.getMessage()), HttpStatus.NOT_FOUND);
        }
        catch (Exception e) {
            log.error("Error loading profile", e);
            return new ResponseEntity<>(new LoadProfileResponse(null, e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @GetMapping("/account/{id}")
    public ResponseEntity<LoadProfileResponse> loadProfileByAccountId(@PathVariable("id") String id) {
        try {
            LoadProfileResponse response = profileService.loadProfileByAccountId(new LoadProfileByAccountIdQuery(
                    UUID.fromString(id)
            ));
            return ResponseEntity.ok(response);
        } catch (ResourceNotFoundException e) {
            return new ResponseEntity<>(new LoadProfileResponse(null, e.getMessage()), HttpStatus.NOT_FOUND);
        }
        catch (Exception e) {
            log.error("Error loading profile", e);
            return new ResponseEntity<>(new LoadProfileResponse(null, e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

		@GetMapping("/link/{link}")
		public ResponseEntity<LoadProfileResponse> loadProfileByLink(@PathVariable("link") String link) {
				try {
						LoadProfileResponse response = profileService.loadProfileByLink(new LoadProfileByLinkQuery(
										link
						));
						return ResponseEntity.ok(response);
				} catch (ResourceNotFoundException e) {
						return new ResponseEntity<>(new LoadProfileResponse(null, e.getMessage()), HttpStatus.NOT_FOUND);
				}
				catch (Exception e) {
						log.error("Error loading profile", e);
						return new ResponseEntity<>(new LoadProfileResponse(null, e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
				}
		}
}
