-- Descripción: cambios realizados para separar mas la logica de negocio por entidades
-- Fecha: 2025-05-10
-- Autor: Diego Axsel Garcia Sierra

-- Modificaciones a la tabla entidad
ALTER TABLE "public".entidad ADD COLUMN logo varchar(255);
COMMENT ON COLUMN "public".entidad.logo IS 'link de minio del logo de la entidad';
UPDATE "public".entidad SET razon_social='platech' WHERE entidad_id=1;
UPDATE "public".entidad SET email='platechoficial' WHERE entidad_id=1;
ALTER TABLE "public".entidad ADD COLUMN colors varchar[];
COMMENT ON COLUMN "public".entidad.colors IS 'colores institucionales';

-- Creación de secuencia y tabla re-recurso-entidad
CREATE SEQUENCE "public"."re_recurso_entidad_seq";
COMMENT ON SEQUENCE "public"."re_recurso_entidad_seq" IS 'secuencia de la tabla re-recurso-entidad';

CREATE TABLE "public"."re_recurso_entidad" (
                                               "re_recurso_entidad_id" bigint DEFAULT nextval('"re_recurso_entidad_seq"'::regclass) NOT NULL,
                                               "activo" boolean,
                                               "entidad_id" bigint,
                                               "recurso_id" bigint,
                                               CONSTRAINT "pk_re_recurso_entidad" PRIMARY KEY ("re_recurso_entidad_id")
);

-- Agregar claves foráneas
ALTER TABLE "public"."re_recurso_entidad"
    ADD CONSTRAINT "fk_re_recurso_entidad_recurso" FOREIGN KEY (recurso_id)
        REFERENCES "public".recurso(recurso_id);

ALTER TABLE "public"."re_recurso_entidad"
    ADD CONSTRAINT "fk_re_recurso_entidad_entidad" FOREIGN KEY (entidad_id)
        REFERENCES "public".entidad(entidad_id);