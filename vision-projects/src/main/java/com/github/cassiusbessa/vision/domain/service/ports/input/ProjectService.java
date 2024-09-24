package com.github.cassiusbessa.vision.domain.service.ports.input;

import com.github.cassiusbessa.vision.domain.service.dtos.project.ProjectDTO;
import com.github.cassiusbessa.vision.domain.service.dtos.project.commands.ProjectCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.project.commands.ProjectDeleteCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.project.commands.ProjectUpdateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.project.queries.LoadProjectsByAccountIdQuery;
import com.github.cassiusbessa.vision.domain.service.dtos.project.responses.LoadedProjectsResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.project.responses.ProjectCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.project.responses.ProjectDeletedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.project.responses.ProjectUpdatedResponse;

import java.util.List;

public interface ProjectService {

    ProjectCreatedResponse createProject(ProjectCreateCommand command);

    ProjectUpdatedResponse updateProject(ProjectUpdateCommand command);

    ProjectDeletedResponse deleteProject(ProjectDeleteCommand command);

    LoadedProjectsResponse loadProjectsByAccountId(LoadProjectsByAccountIdQuery query);
}
