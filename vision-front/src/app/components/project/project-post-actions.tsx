import React, { useEffect, useRef } from 'react';
import { GrSend } from 'react-icons/gr';
import { CiHeart } from 'react-icons/ci';
import { FaRegComment } from 'react-icons/fa6';
import { commentsMock } from '@/app/mocks';
import ProjectPostAddComment from './project-post-add-comment';
import ProjectPostComment from './project-post-comment';

export default function ProjectPostActions({ showAddComment, setShowAddComment }:
{ showAddComment: boolean, setShowAddComment: (showAddComment: boolean) => void }) {
  const textareaRef = useRef<HTMLTextAreaElement>(null);

  const focusTextarea = () => {
    if (textareaRef.current) {
      textareaRef.current.focus();
    }
  };

  useEffect(() => {
    if (showAddComment) {
      focusTextarea();
    }
  }, [showAddComment]);

  return (
    <div className="actions">
      <div className="flex items-center justify-evenly">
        <button type="button" className="btn btn-square btn-ghost w-fit h-fit p-1">
          <CiHeart className="md:w-6 md:h-6 w-8 h-8" />
          <span className="hidden md:block">Curtir</span>
        </button>
        <button
          type="button"
          className="btn btn-square btn-ghost w-fit h-fit p-1"
          onClick={() => setShowAddComment(!showAddComment)}
        >
          <FaRegComment className="md:w-6 md:h-6 w-8 h-8" />
          <span className="hidden md:block">Comentar</span>
        </button>
        <button type="button" className="btn btn-square btn-ghost w-fit h-fit p-1">
          <GrSend className="md:w-6 md:h-6 w-8 h-8" />
          <span className="hidden md:block">Enviar</span>
        </button>
      </div>
      {showAddComment && (
        <>
          <ProjectPostAddComment textareaRef={textareaRef} />
          <ProjectPostComment comment={commentsMock[0]} />
          <ProjectPostComment comment={commentsMock[1]} />
        </>
      )}
    </div>
  );
}
