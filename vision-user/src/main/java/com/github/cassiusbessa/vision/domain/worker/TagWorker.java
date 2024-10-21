package com.github.cassiusbessa.vision.domain.worker;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import com.github.cassiusbessa.vision.domain.service.dtos.tag.CreateTagCommand;
import com.github.cassiusbessa.vision.domain.service.ports.input.TagService;

import jakarta.annotation.PostConstruct;

@Component
public class TagWorker {

	private final TagService tagService;

	@Autowired
	public TagWorker(TagService tagService) {
		this.tagService = tagService;
	}

	@PostConstruct
	public void init() {

		if (tagService.findAll().size() > 0) {
			return;
		}

		startTags();
	}

	public void startTags() {

		List<CreateTagCommand> tags = List.of(
			new CreateTagCommand("Assembly"),
			new CreateTagCommand("C"),
			new CreateTagCommand("C++"),
			new CreateTagCommand("Java"),
			new CreateTagCommand("Python"),
			new CreateTagCommand("Ruby"),
			new CreateTagCommand("JavaScript"),
			new CreateTagCommand("HTML"),
			new CreateTagCommand("CSS"),
			new CreateTagCommand("SQL"),
			new CreateTagCommand("NoSQL"),
			new CreateTagCommand("MongoDB"),
			new CreateTagCommand("PostgreSQL"),
			new CreateTagCommand("React"),
			new CreateTagCommand("Angular"),
			new CreateTagCommand("Vue"),
			new CreateTagCommand("Spring"),
			new CreateTagCommand("Node")
		);

		for (CreateTagCommand tag : tags) {
			tagService.create(tag);
		}
	}

}
