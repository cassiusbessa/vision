package com.github.cassiusbessa.vision.domain.service.mappers;

import org.springframework.stereotype.Component;

import com.github.cassiusbessa.vision.domain.service.dtos.tag.TagDTO;

@Component
public class TagDataMapper {

	public TagDTO tagToTagDTO(com.github.cassiusbessa.vision.domain.core.entities.Tag tag) {
		if (tag == null) {
			return null;
		}
		return new TagDTO(tag.getId().getValue(), tag.getName());
	}
	
}
