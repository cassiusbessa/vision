package com.github.cassiusbessa.vision.http;

import com.github.cassiusbessa.vision.domain.core.exceptions.DomainException;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectUpdateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectUpdatedResponse;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceAlreadyExistsException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceNotFoundException;
import com.github.cassiusbessa.vision.domain.service.exceptions.UnauthorizedException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ValidationException;
import com.github.cassiusbessa.vision.domain.service.ports.input.ProjectService;
import com.github.cassiusbessa.vision.domain.service.ports.input.TokenService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.UUID;

@Slf4j
@RestController
@RequestMapping("/project")
public class ProjectController {

    private final ProjectService projectService;

    private final TokenService tokenService;

    @Autowired
    public ProjectController(ProjectService projectService, TokenService tokenService) {
        this.projectService = projectService;
        this.tokenService = tokenService;
    }

    @PostMapping()
    public ResponseEntity<ProjectCreatedResponse> createProject(@RequestBody ProjectCreateCommand command, @RequestHeader("Authorization") String token) {
        try {
            if (!tokenService.getAccountId(token).equals(command.accountId())) {
                log.error("Unauthorized project creation, invalid token");
                return new ResponseEntity<>(new ProjectCreatedResponse("Unauthorized"), HttpStatus.UNAUTHORIZED);
            }
            ProjectCreatedResponse response = projectService.createProject(command);
            return new ResponseEntity<>(response, HttpStatus.CREATED);
        } catch (ResourceNotFoundException e) {
            return new ResponseEntity<>(new ProjectCreatedResponse(e.getMessage()), HttpStatus.NOT_FOUND);
        } catch (ValidationException | DomainException e) {
            return new ResponseEntity<>(new ProjectCreatedResponse(e.getMessage()), HttpStatus.BAD_REQUEST);
        } catch (ResourceAlreadyExistsException e) {
            return new ResponseEntity<>(new ProjectCreatedResponse(e.getMessage()), HttpStatus.CONFLICT);
        }  catch (UnauthorizedException e) {
            return new ResponseEntity<>(new ProjectCreatedResponse(e.getMessage()), HttpStatus.UNAUTHORIZED);
        } catch (Exception e) {
            log.error("Error creating project", e);
            return new ResponseEntity<>(new ProjectCreatedResponse(e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PutMapping()
    public ResponseEntity<ProjectUpdatedResponse> updateProject(@RequestBody ProjectUpdateCommand command, @RequestHeader("Authorization") String token) {
        try {
            if (!tokenService.getAccountId(token).equals(command.accountId())) {
                log.error("Unauthorized project update, invalid token");
                return new ResponseEntity<>(new ProjectUpdatedResponse("Unauthorized"), HttpStatus.UNAUTHORIZED);
            }
            ProjectUpdatedResponse response = projectService.updateProject(command);
            return new ResponseEntity<>(response, HttpStatus.OK);
        } catch (ResourceNotFoundException e) {
            return new ResponseEntity<>(new ProjectUpdatedResponse(e.getMessage()), HttpStatus.NOT_FOUND);
        } catch (ValidationException | DomainException e) {
            return new ResponseEntity<>(new ProjectUpdatedResponse(e.getMessage()), HttpStatus.BAD_REQUEST);
        } catch (ResourceAlreadyExistsException e) {
            return new ResponseEntity<>(new ProjectUpdatedResponse(e.getMessage()), HttpStatus.CONFLICT);
        }  catch (UnauthorizedException e) {
            return new ResponseEntity<>(new ProjectUpdatedResponse(e.getMessage()), HttpStatus.UNAUTHORIZED);
        } catch (Exception e) {
            log.error("Error creating project", e);
            return new ResponseEntity<>(new ProjectUpdatedResponse(e.getMessage()), HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}
