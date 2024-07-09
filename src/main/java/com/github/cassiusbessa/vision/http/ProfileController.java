package com.github.cassiusbessa.vision.http;

import com.github.cassiusbessa.vision.domain.core.exceptions.DomainException;
import com.github.cassiusbessa.vision.domain.service.dtos.AccountCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.ProfileCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.ProfileCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceAlreadyExistsException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ValidationException;
import com.github.cassiusbessa.vision.domain.service.ports.input.ProfileService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/profile")
public class ProfileController {

    private static final Logger log = LoggerFactory.getLogger(ProfileController.class);
    private final ProfileService profileService;

    @Autowired
    public ProfileController(ProfileService profileService) {
        this.profileService = profileService;
    }

    @PostMapping()
    public ResponseEntity<ProfileCreatedResponse> createProfile(@RequestBody ProfileCreateCommand command) {
        try {
            ProfileCreatedResponse response = profileService.createProfile(command);
            return ResponseEntity.ok(response);
        } catch (ValidationException | DomainException e) {
            return new ResponseEntity<>(new ProfileCreatedResponse(e.getMessage()), HttpStatus.BAD_REQUEST);
        } catch (ResourceAlreadyExistsException e) {
            return new ResponseEntity<>(new ProfileCreatedResponse(e.getMessage()), HttpStatus.CONFLICT);
        } catch (Exception e) {
            log.error("Error creating profile", e);
            return new ResponseEntity<>(new ProfileCreatedResponse(e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

}
