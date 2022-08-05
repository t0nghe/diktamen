import { describe, it, expect } from "vitest";

import { mount } from "@vue/test-utils";
import ArticleItem from "../ArticleItem.vue";

describe("ArticleItem", ()=>{

    it.todo("render article title and info correctly");

    it.todo("complete articles should have classnames that start with complete");

    it.todo("complete articles should display progress circle with goal/goal");

    it.todo("in-progress articles should have classnames that start with progress");
    
    it.todo("in-progress articles should display progress circles with currentIndex/goal");
    
    it.todo("new articles should have classnames that start with new");

    it.todo("new articles should display progress circles with 0/goal");

    it.todo("clicking on a new article, goToArticle event should emitted");

    it.todo("clicking on an in-progress article, goToArticle event should emitted");
})
