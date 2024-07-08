package com.github.cassiusbessa.vision.domain.core.valueobjects;

public class ProjectLinks {

    private final String repository;

    private final String demo;

    public ProjectLinks(String repository, String demo) {
        this.repository = repository;
        this.demo = demo;
    }

    public String getRepository() {
        return repository;
    }

    public String getDemo() {
        return demo;
    }
}
