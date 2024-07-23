package com.github.cassiusbessa.vision.dataaccess.repositories;

import com.github.cassiusbessa.vision.dataaccess.entities.AccountDataBaseEntity;
import com.github.cassiusbessa.vision.dataaccess.entities.ProjectDataBaseEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Repository
@Component
public interface ProjectJpaRepository extends JpaRepository<ProjectDataBaseEntity, UUID> {

    List<ProjectDataBaseEntity> findAllByAccountId(UUID accountId);
}
