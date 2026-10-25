components {
  id: "goal"
  component: "/assets/scripts/goal.script"
}
components {
  id: "wall"
  component: "/assets/scripts/wall.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"grey\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/imagebin/box.tilesource\"\n"
  "}\n"
  ""
  position {
    z: 2.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 1.0\n"
  "restitution: 0.5\n"
  "group: \"wall\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"wall\"\n"
  "  }\n"
  "  data: 0.5\n"
  "  data: 0.5\n"
  "  data: 10.0\n"
  "}\n"
  "event_collision: false\n"
  "event_contact: false\n"
  "event_trigger: false\n"
  ""
}
