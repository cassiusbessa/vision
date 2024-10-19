package com.github.cassiusbessa.vision.http;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.github.cassiusbessa.vision.domain.service.dtos.tag.TagDTO;
import com.github.cassiusbessa.vision.domain.service.ports.input.TagService;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@RestController
@RequestMapping("/tag")
public class TagController {
	private final TagService tagService;

	@Autowired
	public TagController(TagService tagService) {
			this.tagService = tagService;
	}

	@GetMapping()
	public ResponseEntity<List<TagDTO>> findAll() {
		log.info("Finding all tags");
		return ResponseEntity.ok(tagService.findAll());
	}
	
}
