package com.github.cassiusbessa.vision.domain.service.ports.input;

import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreateResponse;

public interface ProjectService {

    ProjectCreateResponse createProject(ProjectCreateCommand command);
}
