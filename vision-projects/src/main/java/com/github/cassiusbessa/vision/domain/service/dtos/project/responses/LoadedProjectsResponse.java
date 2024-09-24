package com.github.cassiusbessa.vision.domain.service.dtos.project.responses;

import com.github.cassiusbessa.vision.domain.service.dtos.project.ProjectDTO;

import java.util.List;

public record LoadedProjectsResponse (List<ProjectDTO> projects, String message) {
}
