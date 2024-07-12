package com.github.cassiusbessa.vision.domain.service.ports.input;

import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreatedResponse;

public interface ProjectService {

    ProjectCreatedResponse createProject(ProjectCreateCommand command);
}
