import React from 'react';

export default function ProjectCardTags({ tags }: { tags: string[] }) {
  return (
    <div className="card-tags h-full flex items-center gap-1">
      {tags.map((tag) => (
        <div key={tag} className="badge badge-outline bg-pink-700">{tag}</div>
      ))}
    </div>
  );
}
