package com.github.cassiusbessa.vision.dataaccess.repositories;

import com.github.cassiusbessa.vision.dataaccess.entities.ProjectDataBaseEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;

@Repository
@Component
public interface ProjectJpaRepository extends JpaRepository<ProjectDataBaseEntity, UUID> {

    List<ProjectDataBaseEntity> findAllByAccountId(UUID accountId);

		@Query(value = "SELECT p.* FROM projects p " +
		"JOIN accounts a ON p.account_id = a.id " +
		"JOIN profiles prof ON prof.account_id = a.id " +
		"WHERE prof.id = :profileId", nativeQuery = true)
List<ProjectDataBaseEntity> findAllByProfileId(@Param("profileId") UUID profileId);

}
