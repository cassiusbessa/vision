package com.github.cassiusbessa.vision.domain.service.ports.input;

import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectUpdateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectUpdatedResponse;

public interface ProjectService {

    ProjectCreatedResponse createProject(ProjectCreateCommand command);

    ProjectUpdatedResponse updateProject(ProjectUpdateCommand command);
}
