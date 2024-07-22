package com.github.cassiusbessa.vision.domain.service.ports.input;

import com.github.cassiusbessa.vision.domain.service.dtos.*;

public interface ProjectService {

    ProjectCreatedResponse createProject(ProjectCreateCommand command);

    ProjectUpdatedResponse updateProject(ProjectUpdateCommand command);

    ProjectDeletedResponse deleteProject(ProjectDeleteCommand command);
}
