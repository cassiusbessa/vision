package com.github.cassiusbessa.vision.domain.service.ports.input;

import java.util.List;

import com.github.cassiusbessa.vision.domain.service.dtos.tag.CreateTagCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.tag.TagDTO;

public interface TagService {
	List<TagDTO> findAll();

	void create(CreateTagCommand command);
}