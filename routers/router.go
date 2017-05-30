// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/api_academica/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/categoria_titulacion",
			beego.NSInclude(
				&controllers.CategoriaTitulacionController{},
			),
		),

		beego.NSNamespace("/condicion_derechos_autor",
			beego.NSInclude(
				&controllers.CondicionDerechosAutorController{},
			),
		),

		beego.NSNamespace("/clase_grupo_academico",
			beego.NSInclude(
				&controllers.ClaseGrupoAcademicoController{},
			),
		),

		beego.NSNamespace("/categoria_criterio_rubrica",
			beego.NSInclude(
				&controllers.CategoriaCriterioRubricaController{},
			),
		),

		beego.NSNamespace("/consejeria",
			beego.NSInclude(
				&controllers.ConsejeriaController{},
			),
		),

		beego.NSNamespace("/cargo_academico",
			beego.NSInclude(
				&controllers.CargoAcademicoController{},
			),
		),

		beego.NSNamespace("/caracteristica_espacio",
			beego.NSInclude(
				&controllers.CaracteristicaEspacioController{},
			),
		),

		beego.NSNamespace("/area_conocimiento",
			beego.NSInclude(
				&controllers.AreaConocimientoController{},
			),
		),

		beego.NSNamespace("/acuerdo",
			beego.NSInclude(
				&controllers.AcuerdoController{},
			),
		),

		beego.NSNamespace("/actividad_recurso_aprendizaje",
			beego.NSInclude(
				&controllers.ActividadRecursoAprendizajeController{},
			),
		),

		beego.NSNamespace("/beneficio_apoyo",
			beego.NSInclude(
				&controllers.BeneficioApoyoController{},
			),
		),

		beego.NSNamespace("/calendario_organizacion",
			beego.NSInclude(
				&controllers.CalendarioOrganizacionController{},
			),
		),

		beego.NSNamespace("/categoria_certificado",
			beego.NSInclude(
				&controllers.CategoriaCertificadoController{},
			),
		),

		beego.NSNamespace("/categoria_ocupacional",
			beego.NSInclude(
				&controllers.CategoriaOcupacionalController{},
			),
		),

		beego.NSNamespace("/categoria_titulo",
			beego.NSInclude(
				&controllers.CategoriaTituloController{},
			),
		),

		beego.NSNamespace("/clasificacion_institucion",
			beego.NSInclude(
				&controllers.ClasificacionInstitucionController{},
			),
		),

		beego.NSNamespace("/espacio_fisico",
			beego.NSInclude(
				&controllers.EspacioFisicoController{},
			),
		),

		beego.NSNamespace("/educacion_distancia",
			beego.NSInclude(
				&controllers.EducacionDistanciaController{},
			),
		),

		beego.NSNamespace("/demografia_estudiante",
			beego.NSInclude(
				&controllers.DemografiaEstudianteController{},
			),
		),

		beego.NSNamespace("/empleado",
			beego.NSInclude(
				&controllers.EmpleadoController{},
			),
		),

		beego.NSNamespace("/empleo_estudiante",
			beego.NSInclude(
				&controllers.EmpleoEstudianteController{},
			),
		),

		beego.NSNamespace("/espacio_academico",
			beego.NSInclude(
				&controllers.EspacioAcademicoController{},
			),
		),

		beego.NSNamespace("/estado_tiempo_matricula",
			beego.NSInclude(
				&controllers.EstadoTiempoMatriculaController{},
			),
		),

		beego.NSNamespace("/estado_empleo",
			beego.NSInclude(
				&controllers.EstadoEmpleoController{},
			),
		),

		beego.NSNamespace("/esquema_identificacion_org",
			beego.NSInclude(
				&controllers.EsquemaIdentificacionOrgController{},
			),
		),

		beego.NSNamespace("/estado_tiempo_completo",
			beego.NSInclude(
				&controllers.EstadoTiempoCompletoController{},
			),
		),

		beego.NSNamespace("/estado_estudiante_proyecto",
			beego.NSInclude(
				&controllers.EstadoEstudianteProyectoController{},
			),
		),

		beego.NSNamespace("/estado_grupo",
			beego.NSInclude(
				&controllers.EstadoGrupoController{},
			),
		),

		beego.NSNamespace("/estado_rol",
			beego.NSInclude(
				&controllers.EstadoRolController{},
			),
		),

		beego.NSNamespace("/estudiante_proyecto_plan_estudios",
			beego.NSInclude(
				&controllers.EstudianteProyectoPlanEstudiosController{},
			),
		),

		beego.NSNamespace("/franja_horaria",
			beego.NSInclude(
				&controllers.FranjaHorariaController{},
			),
		),

		beego.NSNamespace("/modo_acceso",
			beego.NSInclude(
				&controllers.ModoAccesoController{},
			),
		),

		beego.NSNamespace("/metodologia_apoyo",
			beego.NSInclude(
				&controllers.MetodologiaApoyoController{},
			),
		),

		beego.NSNamespace("/metodo_entrada_recurso",
			beego.NSInclude(
				&controllers.MetodoEntradaRecursoController{},
			),
		),

		beego.NSNamespace("/modo_instruccion",
			beego.NSInclude(
				&controllers.ModoInstruccionController{},
			),
		),

		beego.NSNamespace("/matricula_estudiante",
			beego.NSInclude(
				&controllers.MatriculaEstudianteController{},
			),
		),

		beego.NSNamespace("/inscripcion_grupos",
			beego.NSInclude(
				&controllers.InscripcionGruposController{},
			),
		),

		beego.NSNamespace("/limites_creditos",
			beego.NSInclude(
				&controllers.LimitesCreditosController{},
			),
		),

		beego.NSNamespace("/medio_digital",
			beego.NSInclude(
				&controllers.MedioDigitalController{},
			),
		),

		beego.NSNamespace("/metodo_instruccion",
			beego.NSInclude(
				&controllers.MetodoInstruccionController{},
			),
		),

		beego.NSNamespace("/modalidad_programa",
			beego.NSInclude(
				&controllers.ModalidadProgramaController{},
			),
		),

		beego.NSNamespace("/nivel_consideracion_admision",
			beego.NSInclude(
				&controllers.NivelConsideracionAdmisionController{},
			),
		),

		beego.NSNamespace("/nota_estudiante",
			beego.NSInclude(
				&controllers.NotaEstudianteController{},
			),
		),

		beego.NSNamespace("/nivel_semestre",
			beego.NSInclude(
				&controllers.NivelSemestreController{},
			),
		),

		beego.NSNamespace("/oportunidad_aprendizaje",
			beego.NSInclude(
				&controllers.OportunidadAprendizajeController{},
			),
		),

		beego.NSNamespace("/nivel_institucion",
			beego.NSInclude(
				&controllers.NivelInstitucionController{},
			),
		),

		beego.NSNamespace("/nivel_educativo",
			beego.NSInclude(
				&controllers.NivelEducativoController{},
			),
		),

		beego.NSNamespace("/nivel_programa",
			beego.NSInclude(
				&controllers.NivelProgramaController{},
			),
		),

		beego.NSNamespace("/norma_acuerdo",
			beego.NSInclude(
				&controllers.NormaAcuerdoController{},
			),
		),

		beego.NSNamespace("/nucleo_basico_conocimiento",
			beego.NSInclude(
				&controllers.NucleoBasicoConocimientoController{},
			),
		),

		beego.NSNamespace("/proposito_recurso",
			beego.NSInclude(
				&controllers.PropositoRecursoController{},
			),
		),

		beego.NSNamespace("/persona",
			beego.NSInclude(
				&controllers.PersonaController{},
			),
		),

		beego.NSNamespace("/perfil_titulacion",
			beego.NSInclude(
				&controllers.PerfilTitulacionController{},
			),
		),

		beego.NSNamespace("/preinscripcion_espacios",
			beego.NSInclude(
				&controllers.PreinscripcionEspaciosController{},
			),
		),

		beego.NSNamespace("/prioridad_consejeria",
			beego.NSInclude(
				&controllers.PrioridadConsejeriaController{},
			),
		),

		beego.NSNamespace("/proyecto_curricular",
			beego.NSInclude(
				&controllers.ProyectoCurricularController{},
			),
		),

		beego.NSNamespace("/razon_salida_programa",
			beego.NSInclude(
				&controllers.RazonSalidaProgramaController{},
			),
		),

		beego.NSNamespace("/ps_institucion",
			beego.NSInclude(
				&controllers.PsInstitucionController{},
			),
		),

		beego.NSNamespace("/ref_norma_acuerdo",
			beego.NSInclude(
				&controllers.RefNormaAcuerdoController{},
			),
		),

		beego.NSNamespace("/responsable",
			beego.NSInclude(
				&controllers.ResponsableController{},
			),
		),

		beego.NSNamespace("/prueba_admision_estudiante",
			beego.NSInclude(
				&controllers.PruebaAdmisionEstudianteController{},
			),
		),

		beego.NSNamespace("/razon_separacion",
			beego.NSInclude(
				&controllers.RazonSeparacionController{},
			),
		),

		beego.NSNamespace("/ref_estado_estudiante_proyecto",
			beego.NSInclude(
				&controllers.RefEstadoEstudianteProyectoController{},
			),
		),

		beego.NSNamespace("/ref_estado_grupo",
			beego.NSInclude(
				&controllers.RefEstadoGrupoController{},
			),
		),

		beego.NSNamespace("/ref_tipo_relacion_org",
			beego.NSInclude(
				&controllers.RefTipoRelacionOrgController{},
			),
		),

		beego.NSNamespace("/tipo_api_acceso",
			beego.NSInclude(
				&controllers.TipoApiAccesoController{},
			),
		),

		beego.NSNamespace("/soporte_fisico",
			beego.NSInclude(
				&controllers.SoporteFisicoController{},
			),
		),

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

		beego.NSNamespace("/tipo_espacio_academico",
			beego.NSInclude(
				&controllers.TipoEspacioAcademicoController{},
			),
		),

		beego.NSNamespace("/rubrica",
			beego.NSInclude(
				&controllers.RubricaController{},
			),
		),

		beego.NSNamespace("/rol",
			beego.NSInclude(
				&controllers.RolController{},
			),
		),

		beego.NSNamespace("/sistema_calendario",
			beego.NSInclude(
				&controllers.SistemaCalendarioController{},
			),
		),

		beego.NSNamespace("/solicitud_apoyo",
			beego.NSInclude(
				&controllers.SolicitudApoyoController{},
			),
		),

		beego.NSNamespace("/solicitud_consejeria",
			beego.NSInclude(
				&controllers.SolicitudConsejeriaController{},
			),
		),

		beego.NSNamespace("/tipo_admision",
			beego.NSInclude(
				&controllers.TipoAdmisionController{},
			),
		),

		beego.NSNamespace("/tipo_admision_institucion",
			beego.NSInclude(
				&controllers.TipoAdmisionInstitucionController{},
			),
		),

		beego.NSNamespace("/tipo_autor",
			beego.NSInclude(
				&controllers.TipoAutorController{},
			),
		),

		beego.NSNamespace("/tipo_jornada",
			beego.NSInclude(
				&controllers.TipoJornadaController{},
			),
		),

		beego.NSNamespace("/tipo_observacion_grupo",
			beego.NSInclude(
				&controllers.TipoObservacionGrupoController{},
			),
		),

		beego.NSNamespace("/tipo_relacion_org",
			beego.NSInclude(
				&controllers.TipoRelacionOrgController{},
			),
		),

		beego.NSNamespace("/tipo_riesgo_recurso",
			beego.NSInclude(
				&controllers.TipoRiesgoRecursoController{},
			),
		),

		beego.NSNamespace("/tipo_identificacion_org",
			beego.NSInclude(
				&controllers.TipoIdentificacionOrgController{},
			),
		),

		beego.NSNamespace("/tipo_observacion_nota",
			beego.NSInclude(
				&controllers.TipoObservacionNotaController{},
			),
		),

		beego.NSNamespace("/tipo_participacion",
			beego.NSInclude(
				&controllers.TipoParticipacionController{},
			),
		),

		beego.NSNamespace("/tipo_prueba_admision",
			beego.NSInclude(
				&controllers.TipoPruebaAdmisionController{},
			),
		),

		beego.NSNamespace("/titulo_otorgado",
			beego.NSInclude(
				&controllers.TituloOtorgadoController{},
			),
		),

		beego.NSNamespace("/tipo_evento",
			beego.NSInclude(
				&controllers.TipoEventoController{},
			),
		),

		beego.NSNamespace("/tipo_interactividad",
			beego.NSInclude(
				&controllers.TipoInteractividadController{},
			),
		),

		beego.NSNamespace("/tipo_matricula",
			beego.NSInclude(
				&controllers.TipoMatriculaController{},
			),
		),

		beego.NSNamespace("/tipo_organizacion",
			beego.NSInclude(
				&controllers.TipoOrganizacionController{},
			),
		),

		beego.NSNamespace("/tipo_programa_laboral",
			beego.NSInclude(
				&controllers.TipoProgramaLaboralController{},
			),
		),

		beego.NSNamespace("/tipo_recurso_aprendizaje",
			beego.NSInclude(
				&controllers.TipoRecursoAprendizajeController{},
			),
		),

		beego.NSNamespace("/tipo_separacion",
			beego.NSInclude(
				&controllers.TipoSeparacionController{},
			),
		),

		beego.NSNamespace("/tipo_solicitud_apoyo",
			beego.NSInclude(
				&controllers.TipoSolicitudApoyoController{},
			),
		),

		beego.NSNamespace("/tipo_solicitud_consejeria",
			beego.NSInclude(
				&controllers.TipoSolicitudConsejeriaController{},
			),
		),

		beego.NSNamespace("/transferencia_lista",
			beego.NSInclude(
				&controllers.TransferenciaListaController{},
			),
		),

		beego.NSNamespace("/tipo_residencia",
			beego.NSInclude(
				&controllers.TipoResidenciaController{},
			),
		),

		beego.NSNamespace("/razon_fin_grupo_docente",
			beego.NSInclude(
				&controllers.RazonFinGrupoDocenteController{},
			),
		),

		beego.NSNamespace("/control_institucion",
			beego.NSInclude(
				&controllers.ControlInstitucionController{},
			),
		),

		beego.NSNamespace("/contenido_prog_docente",
			beego.NSInclude(
				&controllers.ContenidoProgDocenteController{},
			),
		),

		beego.NSNamespace("/corte_evaluacion",
			beego.NSInclude(
				&controllers.CorteEvaluacionController{},
			),
		),

		beego.NSNamespace("/grupo_academico",
			beego.NSInclude(
				&controllers.GrupoAcademicoController{},
			),
		),

		beego.NSNamespace("/grupo_actividad",
			beego.NSInclude(
				&controllers.GrupoActividadController{},
			),
		),

		beego.NSNamespace("/asignacion_docente_clase",
			beego.NSInclude(
				&controllers.AsignacionDocenteClaseController{},
			),
		),

		beego.NSNamespace("/organizacion_persona_rol",
			beego.NSInclude(
				&controllers.OrganizacionPersonaRolController{},
			),
		),

		beego.NSNamespace("/horario",
			beego.NSInclude(
				&controllers.HorarioController{},
			),
		),

		beego.NSNamespace("/estado_dependencia",
			beego.NSInclude(
				&controllers.EstadoDependenciaController{},
			),
		),

		beego.NSNamespace("/programa_laboral",
			beego.NSInclude(
				&controllers.ProgramaLaboralController{},
			),
		),

		beego.NSNamespace("/organizacion",
			beego.NSInclude(
				&controllers.OrganizacionController{},
			),
		),

		beego.NSNamespace("/tipo_residencia_campus",
			beego.NSInclude(
				&controllers.TipoResidenciaCampusController{},
			),
		),

		beego.NSNamespace("/espacio_academico_pregrado",
			beego.NSInclude(
				&controllers.EspacioAcademicoPregradoController{},
			),
		),

		beego.NSNamespace("/sesion_organizacion",
			beego.NSInclude(
				&controllers.SesionOrganizacionController{},
			),
		),

		beego.NSNamespace("/estudiante_proyecto",
			beego.NSInclude(
				&controllers.EstudianteProyectoController{},
			),
		),

		beego.NSNamespace("/apoyo_financiero",
			beego.NSInclude(
				&controllers.ApoyoFinancieroController{},
			),
		),

		beego.NSNamespace("/soporte_fisico_recurso",
			beego.NSInclude(
				&controllers.SoporteFisicoRecursoController{},
			),
		),

		beego.NSNamespace("/solicitud_admision",
			beego.NSInclude(
				&controllers.SolicitudAdmisionController{},
			),
		),

		beego.NSNamespace("/proyecto_curricular_maestria",
			beego.NSInclude(
				&controllers.ProyectoCurricularMaestriaController{},
			),
		),

		beego.NSNamespace("/unidad_dedicacion",
			beego.NSInclude(
				&controllers.UnidadDedicacionController{},
			),
		),

		beego.NSNamespace("/actividad_aprendizaje",
			beego.NSInclude(
				&controllers.ActividadAprendizajeController{},
			),
		),

		beego.NSNamespace("/plan_estudios",
			beego.NSInclude(
				&controllers.PlanEstudiosController{},
			),
		),

		beego.NSNamespace("/criterio_rubrica",
			beego.NSInclude(
				&controllers.CriterioRubricaController{},
			),
		),

		beego.NSNamespace("/contenido_programatico",
			beego.NSInclude(
				&controllers.ContenidoProgramaticoController{},
			),
		),

		beego.NSNamespace("/empleado_postsecundaria",
			beego.NSInclude(
				&controllers.EmpleadoPostsecundariaController{},
			),
		),

		beego.NSNamespace("/dia_asesoria_docente",
			beego.NSInclude(
				&controllers.DiaAsesoriaDocenteController{},
			),
		),

		beego.NSNamespace("/recurso_aprendizaje",
			beego.NSInclude(
				&controllers.RecursoAprendizajeController{},
			),
		),

		beego.NSNamespace("/grupo_plan_estudios",
			beego.NSInclude(
				&controllers.GrupoPlanEstudiosController{},
			),
		),

		beego.NSNamespace("/espacios_homologables",
			beego.NSInclude(
				&controllers.EspaciosHomologablesController{},
			),
		),

		beego.NSNamespace("/evento_organizacion",
			beego.NSInclude(
				&controllers.EventoOrganizacionController{},
			),
		),

		beego.NSNamespace("/espacios_prerrequisitos",
			beego.NSInclude(
				&controllers.EspaciosPrerrequisitosController{},
			),
		),

		beego.NSNamespace("/formato_libro",
			beego.NSInclude(
				&controllers.FormatoLibroController{},
			),
		),

		beego.NSNamespace("/inclusion_promedio",
			beego.NSInclude(
				&controllers.InclusionPromedioController{},
			),
		),

		beego.NSNamespace("/identificador_organizacion",
			beego.NSInclude(
				&controllers.IdentificadorOrganizacionController{},
			),
		),

		beego.NSNamespace("/idioma",
			beego.NSInclude(
				&controllers.IdiomaController{},
			),
		),

		beego.NSNamespace("/nivel_criterio_rubrica",
			beego.NSInclude(
				&controllers.NivelCriterioRubricaController{},
			),
		),

		beego.NSNamespace("/participacion_programa",
			beego.NSInclude(
				&controllers.ParticipacionProgramaController{},
			),
		),

		beego.NSNamespace("/nivel_educativo_recurso",
			beego.NSInclude(
				&controllers.NivelEducativoRecursoController{},
			),
		),

		beego.NSNamespace("/responsable_sesion",
			beego.NSInclude(
				&controllers.ResponsableSesionController{},
			),
		),

		beego.NSNamespace("/plan_estudios_espacio",
			beego.NSInclude(
				&controllers.PlanEstudiosEspacioController{},
			),
		),

		beego.NSNamespace("/proyecto_curricular_pregrado",
			beego.NSInclude(
				&controllers.ProyectoCurricularPregradoController{},
			),
		),

		beego.NSNamespace("/resumen_academico",
			beego.NSInclude(
				&controllers.ResumenAcademicoController{},
			),
		),

		beego.NSNamespace("/tipo_estado_rol",
			beego.NSInclude(
				&controllers.TipoEstadoRolController{},
			),
		),

		beego.NSNamespace("/tipo_actividad",
			beego.NSInclude(
				&controllers.TipoActividadController{},
			),
		),

		beego.NSNamespace("/unidad_tiempo",
			beego.NSInclude(
				&controllers.UnidadTiempoController{},
			),
		),

		beego.NSNamespace("/tipo_sesion",
			beego.NSInclude(
				&controllers.TipoSesionController{},
			),
		),

		beego.NSNamespace("/usuario_final",
			beego.NSInclude(
				&controllers.UsuarioFinalController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
