import React from 'react';

export default function ProjectCardTags({ tags }: { tags: { id: string, name:string }[] }) {
  return (
    <div className="card-tags h-full flex items-center gap-1">
      {tags.map((tag) => (
        <div key={tag.id} className="badge badge-outline bg-pink-700">{tag.name}</div>
      ))}
    </div>
  );
}
