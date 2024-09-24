package com.github.cassiusbessa.vision.dataaccess;

import com.github.cassiusbessa.vision.dataaccess.entities.TagDataBaseEntity;
import com.github.cassiusbessa.vision.dataaccess.mappers.TagDataBaseMapper;
import com.github.cassiusbessa.vision.dataaccess.repositories.TagJpaRepository;
import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.service.ports.output.TagRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;

@Repository
public class TagRepositoryImpl implements TagRepository {

    private final TagJpaRepository tagRepository;
    private final TagDataBaseMapper tagDataBaseMapper;

    @Autowired
    public TagRepositoryImpl(TagJpaRepository tagRepository, TagDataBaseMapper tagDataBaseMapper) {
        this.tagRepository = tagRepository;
        this.tagDataBaseMapper = tagDataBaseMapper;
    }

    @Override
    public boolean existsAllTags(List<UUID> tagIds) {
        return false;
    }

    @Override
    public List<Tag> findAllById(List<UUID> tagIds) {

        List<TagDataBaseEntity> tagDataBaseEntities = tagRepository.findAllById(tagIds);
        return tagDataBaseEntities.stream().map(tagDataBaseMapper::dbEntityToTag).toList();
    }
}
