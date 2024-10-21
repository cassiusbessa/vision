package com.github.cassiusbessa.vision.domain.service;

import java.util.List;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.core.valueobjects.TagId;
import com.github.cassiusbessa.vision.domain.service.dtos.tag.CreateTagCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.tag.TagDTO;
import com.github.cassiusbessa.vision.domain.service.mappers.TagDataMapper;
import com.github.cassiusbessa.vision.domain.service.ports.input.TagService;
import com.github.cassiusbessa.vision.domain.service.ports.output.TagRepository;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@Service
public class TagServiceImpl implements TagService {


	private final TagRepository tagRepository;

	private final TagDataMapper tagDataMapper;

	@Autowired
	public TagServiceImpl(TagRepository tagRepository, TagDataMapper tagDataMapper) {
		this.tagRepository = tagRepository;
		this.tagDataMapper = tagDataMapper;
	}


	@Override
	public List<TagDTO> findAll() {
		log.info("Finding all tags");
		return tagRepository.findAll().stream().map(tagDataMapper::tagToTagDTO).toList();
	}

	@Override
	public void create(CreateTagCommand command) {
		log.info("Creating tag with name: {}", command.name());


		TagId tagId = new TagId(UUID.randomUUID());
		Tag tag = new Tag(tagId, command.name());

		tagRepository.create(tag);
	}
	
}
