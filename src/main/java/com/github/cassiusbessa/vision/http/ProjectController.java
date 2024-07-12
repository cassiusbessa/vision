package com.github.cassiusbessa.vision.http;

import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.ports.input.ProjectService;
import com.github.cassiusbessa.vision.domain.service.ports.input.TokenService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

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
    public ResponseEntity<ProjectCreatedResponse> createProject(@RequestBody ProjectCreateCommand command) {
        try {
            ProjectCreatedResponse response = projectService.createProject(command);
            return ResponseEntity.ok(response);
        } catch (Exception e) {
            log.error("Error creating project", e);
            return ResponseEntity.badRequest().build();
        }
    }
}
