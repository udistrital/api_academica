package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["api_academica/controllers:ActividadAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ActividadAprendizajeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ActividadAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ActividadAprendizajeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ActividadAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ActividadAprendizajeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ActividadAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ActividadAprendizajeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ActividadAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ActividadAprendizajeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ActividadRecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ActividadRecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ActividadRecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ActividadRecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ActividadRecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ActividadRecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ActividadRecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ActividadRecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ActividadRecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ActividadRecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AcuerdoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AcuerdoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AcuerdoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AcuerdoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AcuerdoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ApoyoFinancieroController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ApoyoFinancieroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ApoyoFinancieroController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ApoyoFinancieroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ApoyoFinancieroController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ApoyoFinancieroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ApoyoFinancieroController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ApoyoFinancieroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ApoyoFinancieroController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ApoyoFinancieroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AsignacionDocenteClaseController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AsignacionDocenteClaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AsignacionDocenteClaseController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AsignacionDocenteClaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AsignacionDocenteClaseController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AsignacionDocenteClaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AsignacionDocenteClaseController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AsignacionDocenteClaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:AsignacionDocenteClaseController"] = append(beego.GlobalControllerRouter["api_academica/controllers:AsignacionDocenteClaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:BeneficioApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:BeneficioApoyoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:BeneficioApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:BeneficioApoyoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:BeneficioApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:BeneficioApoyoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:BeneficioApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:BeneficioApoyoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:BeneficioApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:BeneficioApoyoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CalendarioOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CalendarioOrganizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CalendarioOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CalendarioOrganizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CalendarioOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CalendarioOrganizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CalendarioOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CalendarioOrganizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CalendarioOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CalendarioOrganizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CaracteristicaEspacioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CaracteristicaEspacioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CaracteristicaEspacioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CaracteristicaEspacioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CaracteristicaEspacioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CaracteristicaEspacioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CaracteristicaEspacioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CaracteristicaEspacioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CaracteristicaEspacioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CaracteristicaEspacioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CargoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CargoAcademicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CargoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CargoAcademicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CargoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CargoAcademicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CargoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CargoAcademicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CargoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CargoAcademicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaCertificadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaCertificadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaCertificadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaCertificadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaCertificadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaCertificadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaCertificadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaCertificadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaCertificadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaCertificadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaCriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaCriterioRubricaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaCriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaCriterioRubricaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaCriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaCriterioRubricaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaCriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaCriterioRubricaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaCriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaCriterioRubricaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaOcupacionalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaOcupacionalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaOcupacionalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaOcupacionalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaOcupacionalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaOcupacionalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaOcupacionalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaOcupacionalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaOcupacionalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaOcupacionalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaTitulacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaTitulacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaTitulacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaTitulacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaTitulacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaTitulacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaTitulacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaTitulacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaTitulacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaTitulacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaTituloController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaTituloController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaTituloController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaTituloController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaTituloController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaTituloController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaTituloController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaTituloController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CategoriaTituloController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CategoriaTituloController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ClaseGrupoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ClaseGrupoAcademicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ClaseGrupoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ClaseGrupoAcademicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ClaseGrupoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ClaseGrupoAcademicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ClaseGrupoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ClaseGrupoAcademicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ClaseGrupoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ClaseGrupoAcademicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ClasificacionInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ClasificacionInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ClasificacionInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ClasificacionInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ClasificacionInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ClasificacionInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ClasificacionInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ClasificacionInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ClasificacionInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ClasificacionInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CondicionDerechosAutorController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CondicionDerechosAutorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CondicionDerechosAutorController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CondicionDerechosAutorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CondicionDerechosAutorController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CondicionDerechosAutorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CondicionDerechosAutorController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CondicionDerechosAutorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CondicionDerechosAutorController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CondicionDerechosAutorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ConsejeriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ConsejeriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ConsejeriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ConsejeriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ConsejeriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgramaticoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgramaticoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgramaticoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgramaticoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgramaticoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgramaticoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgramaticoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgramaticoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgramaticoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ContenidoProgramaticoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ControlInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ControlInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ControlInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ControlInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ControlInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ControlInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ControlInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ControlInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ControlInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ControlInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CorteEvaluacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CorteEvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CorteEvaluacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CorteEvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CorteEvaluacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CorteEvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CorteEvaluacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CorteEvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CorteEvaluacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CorteEvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CriterioRubricaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CriterioRubricaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CriterioRubricaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CriterioRubricaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:CriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:CriterioRubricaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:DemografiaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:DemografiaEstudianteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:DemografiaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:DemografiaEstudianteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:DemografiaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:DemografiaEstudianteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:DemografiaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:DemografiaEstudianteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:DemografiaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:DemografiaEstudianteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:DiaAsesoriaDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:DiaAsesoriaDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:DiaAsesoriaDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:DiaAsesoriaDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:DiaAsesoriaDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:DiaAsesoriaDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:DiaAsesoriaDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:DiaAsesoriaDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:DiaAsesoriaDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:DiaAsesoriaDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EducacionDistanciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EducacionDistanciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EducacionDistanciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EducacionDistanciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EducacionDistanciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EducacionDistanciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EducacionDistanciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EducacionDistanciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EducacionDistanciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EducacionDistanciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleadoPostsecundariaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleadoPostsecundariaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleadoPostsecundariaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleadoPostsecundariaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleadoPostsecundariaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleadoPostsecundariaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleadoPostsecundariaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleadoPostsecundariaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleadoPostsecundariaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleadoPostsecundariaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleoEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleoEstudianteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleoEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleoEstudianteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleoEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleoEstudianteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleoEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleoEstudianteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EmpleoEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EmpleoEstudianteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoPregradoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoPregradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoPregradoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoPregradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoPregradoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoPregradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoPregradoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoPregradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoPregradoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioAcademicoPregradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspaciosHomologablesController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspaciosHomologablesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspaciosHomologablesController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspaciosHomologablesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspaciosHomologablesController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspaciosHomologablesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspaciosHomologablesController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspaciosHomologablesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspaciosHomologablesController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspaciosHomologablesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspaciosPrerrequisitosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspaciosPrerrequisitosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspaciosPrerrequisitosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspaciosPrerrequisitosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspaciosPrerrequisitosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspaciosPrerrequisitosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspaciosPrerrequisitosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspaciosPrerrequisitosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EspaciosPrerrequisitosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EspaciosPrerrequisitosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EsquemaIdentificacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EsquemaIdentificacionOrgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EsquemaIdentificacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EsquemaIdentificacionOrgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EsquemaIdentificacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EsquemaIdentificacionOrgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EsquemaIdentificacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EsquemaIdentificacionOrgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EsquemaIdentificacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EsquemaIdentificacionOrgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoDependenciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoDependenciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoDependenciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoDependenciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoDependenciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoEmpleoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoEmpleoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoEmpleoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoEmpleoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoEmpleoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoEmpleoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoEmpleoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoEmpleoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoEmpleoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoEmpleoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoEstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoEstudianteProyectoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoEstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoEstudianteProyectoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoEstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoEstudianteProyectoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoEstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoEstudianteProyectoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoEstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoEstudianteProyectoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoGrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoGrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoGrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoGrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoGrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoRolController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoRolController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoRolController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoRolController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoRolController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoCompletoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoCompletoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoCompletoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoCompletoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoCompletoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoCompletoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoCompletoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoCompletoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoCompletoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoCompletoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoMatriculaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoMatriculaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoMatriculaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoMatriculaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoMatriculaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoMatriculaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoMatriculaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoMatriculaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoMatriculaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstadoTiempoMatriculaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoPlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoPlanEstudiosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoPlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoPlanEstudiosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoPlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoPlanEstudiosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoPlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoPlanEstudiosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoPlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EstudianteProyectoPlanEstudiosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EventoOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EventoOrganizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EventoOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EventoOrganizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EventoOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EventoOrganizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EventoOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EventoOrganizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:EventoOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:EventoOrganizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:FormatoLibroController"] = append(beego.GlobalControllerRouter["api_academica/controllers:FormatoLibroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:FormatoLibroController"] = append(beego.GlobalControllerRouter["api_academica/controllers:FormatoLibroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:FormatoLibroController"] = append(beego.GlobalControllerRouter["api_academica/controllers:FormatoLibroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:FormatoLibroController"] = append(beego.GlobalControllerRouter["api_academica/controllers:FormatoLibroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:FormatoLibroController"] = append(beego.GlobalControllerRouter["api_academica/controllers:FormatoLibroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:FranjaHorariaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:FranjaHorariaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:FranjaHorariaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:FranjaHorariaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:FranjaHorariaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:FranjaHorariaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:FranjaHorariaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:FranjaHorariaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:FranjaHorariaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:FranjaHorariaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoAcademicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoAcademicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoAcademicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoAcademicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoAcademicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoActividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoActividadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoActividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoActividadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoActividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoActividadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoActividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoActividadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoActividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoActividadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoPlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoPlanEstudiosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoPlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoPlanEstudiosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoPlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoPlanEstudiosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoPlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoPlanEstudiosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:GrupoPlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:GrupoPlanEstudiosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:HorarioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:HorarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:HorarioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:HorarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:HorarioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:HorarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:HorarioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:HorarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:HorarioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:HorarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:IdentificadorOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:IdentificadorOrganizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:IdentificadorOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:IdentificadorOrganizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:IdentificadorOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:IdentificadorOrganizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:IdentificadorOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:IdentificadorOrganizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:IdentificadorOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:IdentificadorOrganizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:IdiomaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:IdiomaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:IdiomaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:IdiomaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:IdiomaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:IdiomaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:InclusionPromedioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:InclusionPromedioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:InclusionPromedioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:InclusionPromedioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:InclusionPromedioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:InclusionPromedioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:InclusionPromedioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:InclusionPromedioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:InclusionPromedioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:InclusionPromedioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:InscripcionGruposController"] = append(beego.GlobalControllerRouter["api_academica/controllers:InscripcionGruposController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:InscripcionGruposController"] = append(beego.GlobalControllerRouter["api_academica/controllers:InscripcionGruposController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:InscripcionGruposController"] = append(beego.GlobalControllerRouter["api_academica/controllers:InscripcionGruposController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:InscripcionGruposController"] = append(beego.GlobalControllerRouter["api_academica/controllers:InscripcionGruposController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:InscripcionGruposController"] = append(beego.GlobalControllerRouter["api_academica/controllers:InscripcionGruposController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:LimitesCreditosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:LimitesCreditosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:LimitesCreditosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:LimitesCreditosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:LimitesCreditosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:LimitesCreditosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:LimitesCreditosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:LimitesCreditosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:LimitesCreditosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:LimitesCreditosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MatriculaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MatriculaEstudianteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MatriculaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MatriculaEstudianteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MatriculaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MatriculaEstudianteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MatriculaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MatriculaEstudianteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MatriculaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MatriculaEstudianteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MedioDigitalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MedioDigitalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MedioDigitalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MedioDigitalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MedioDigitalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MedioDigitalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MedioDigitalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MedioDigitalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MedioDigitalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MedioDigitalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodoEntradaRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodoEntradaRecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodoEntradaRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodoEntradaRecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodoEntradaRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodoEntradaRecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodoEntradaRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodoEntradaRecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodoEntradaRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodoEntradaRecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodoInstruccionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodoInstruccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodoInstruccionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodoInstruccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodoInstruccionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodoInstruccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodoInstruccionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodoInstruccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodoInstruccionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodoInstruccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodologiaApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodologiaApoyoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodologiaApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodologiaApoyoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodologiaApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodologiaApoyoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodologiaApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodologiaApoyoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:MetodologiaApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:MetodologiaApoyoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModalidadProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModalidadProgramaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModalidadProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModalidadProgramaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModalidadProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModalidadProgramaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModalidadProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModalidadProgramaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModalidadProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModalidadProgramaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModoAccesoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModoAccesoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModoAccesoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModoAccesoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModoAccesoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModoAccesoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModoAccesoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModoAccesoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModoAccesoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModoAccesoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModoInstruccionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModoInstruccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModoInstruccionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModoInstruccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModoInstruccionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModoInstruccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModoInstruccionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModoInstruccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ModoInstruccionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ModoInstruccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelConsideracionAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelConsideracionAdmisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelConsideracionAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelConsideracionAdmisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelConsideracionAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelConsideracionAdmisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelConsideracionAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelConsideracionAdmisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelConsideracionAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelConsideracionAdmisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelCriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelCriterioRubricaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelCriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelCriterioRubricaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelCriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelCriterioRubricaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelCriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelCriterioRubricaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelCriterioRubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelCriterioRubricaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoRecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoRecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoRecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoRecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelEducativoRecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelProgramaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelProgramaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelProgramaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelProgramaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelProgramaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelSemestreController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelSemestreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelSemestreController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelSemestreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelSemestreController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelSemestreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelSemestreController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelSemestreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NivelSemestreController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NivelSemestreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NormaAcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NormaAcuerdoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NormaAcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NormaAcuerdoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NormaAcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NormaAcuerdoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NormaAcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NormaAcuerdoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NormaAcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NormaAcuerdoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NotaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NotaEstudianteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NotaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NotaEstudianteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NotaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NotaEstudianteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NotaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NotaEstudianteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NotaEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NotaEstudianteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NucleoBasicoConocimientoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NucleoBasicoConocimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NucleoBasicoConocimientoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NucleoBasicoConocimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NucleoBasicoConocimientoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NucleoBasicoConocimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NucleoBasicoConocimientoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NucleoBasicoConocimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:NucleoBasicoConocimientoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:NucleoBasicoConocimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OportunidadAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OportunidadAprendizajeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OportunidadAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OportunidadAprendizajeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OportunidadAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OportunidadAprendizajeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OportunidadAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OportunidadAprendizajeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OportunidadAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OportunidadAprendizajeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OrganizacionPersonaRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OrganizacionPersonaRolController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OrganizacionPersonaRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OrganizacionPersonaRolController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OrganizacionPersonaRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OrganizacionPersonaRolController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OrganizacionPersonaRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OrganizacionPersonaRolController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:OrganizacionPersonaRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:OrganizacionPersonaRolController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ParticipacionProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ParticipacionProgramaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ParticipacionProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ParticipacionProgramaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ParticipacionProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ParticipacionProgramaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ParticipacionProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ParticipacionProgramaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ParticipacionProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ParticipacionProgramaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PerfilTitulacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PerfilTitulacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PerfilTitulacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PerfilTitulacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PerfilTitulacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PerfilTitulacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PerfilTitulacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PerfilTitulacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PerfilTitulacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PerfilTitulacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PersonaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PersonaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PersonaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PersonaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PersonaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosEspacioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosEspacioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosEspacioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosEspacioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosEspacioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosEspacioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosEspacioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosEspacioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosEspacioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PlanEstudiosEspacioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PreinscripcionEspaciosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PreinscripcionEspaciosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PreinscripcionEspaciosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PreinscripcionEspaciosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PreinscripcionEspaciosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PreinscripcionEspaciosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PreinscripcionEspaciosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PreinscripcionEspaciosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PreinscripcionEspaciosController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PreinscripcionEspaciosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PrioridadConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PrioridadConsejeriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PrioridadConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PrioridadConsejeriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PrioridadConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PrioridadConsejeriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PrioridadConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PrioridadConsejeriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PrioridadConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PrioridadConsejeriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProgramaLaboralController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProgramaLaboralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProgramaLaboralController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProgramaLaboralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProgramaLaboralController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProgramaLaboralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProgramaLaboralController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProgramaLaboralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProgramaLaboralController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProgramaLaboralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PropositoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PropositoRecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PropositoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PropositoRecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PropositoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PropositoRecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PropositoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PropositoRecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PropositoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PropositoRecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularMaestriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularMaestriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularMaestriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularMaestriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularMaestriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularMaestriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularMaestriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularMaestriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularMaestriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularMaestriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularPregradoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularPregradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularPregradoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularPregradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularPregradoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularPregradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularPregradoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularPregradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularPregradoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ProyectoCurricularPregradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PruebaAdmisionEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PruebaAdmisionEstudianteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PruebaAdmisionEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PruebaAdmisionEstudianteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PruebaAdmisionEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PruebaAdmisionEstudianteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PruebaAdmisionEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PruebaAdmisionEstudianteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PruebaAdmisionEstudianteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PruebaAdmisionEstudianteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PsInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PsInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PsInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PsInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PsInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PsInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PsInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PsInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:PsInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:PsInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonFinGrupoDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonFinGrupoDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonFinGrupoDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonFinGrupoDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonFinGrupoDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonFinGrupoDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonFinGrupoDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonFinGrupoDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonFinGrupoDocenteController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonFinGrupoDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonSalidaProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonSalidaProgramaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonSalidaProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonSalidaProgramaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonSalidaProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonSalidaProgramaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonSalidaProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonSalidaProgramaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonSalidaProgramaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonSalidaProgramaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonSeparacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonSeparacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonSeparacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonSeparacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonSeparacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonSeparacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonSeparacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonSeparacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RazonSeparacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RazonSeparacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefEstadoEstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefEstadoEstudianteProyectoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefEstadoEstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefEstadoEstudianteProyectoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefEstadoEstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefEstadoEstudianteProyectoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefEstadoEstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefEstadoEstudianteProyectoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefEstadoEstudianteProyectoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefEstadoEstudianteProyectoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefEstadoGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefEstadoGrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefEstadoGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefEstadoGrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefEstadoGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefEstadoGrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefEstadoGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefEstadoGrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefEstadoGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefEstadoGrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefNormaAcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefNormaAcuerdoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefNormaAcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefNormaAcuerdoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefNormaAcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefNormaAcuerdoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefNormaAcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefNormaAcuerdoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefNormaAcuerdoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefNormaAcuerdoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefTipoRelacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefTipoRelacionOrgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefTipoRelacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefTipoRelacionOrgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefTipoRelacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefTipoRelacionOrgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefTipoRelacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefTipoRelacionOrgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RefTipoRelacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RefTipoRelacionOrgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResponsableController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResponsableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResponsableController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResponsableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResponsableController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResponsableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResponsableController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResponsableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResponsableController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResponsableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResponsableSesionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResponsableSesionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResponsableSesionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResponsableSesionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResponsableSesionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResponsableSesionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResponsableSesionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResponsableSesionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResponsableSesionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResponsableSesionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResumenAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResumenAcademicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResumenAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResumenAcademicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResumenAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResumenAcademicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResumenAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResumenAcademicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:ResumenAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:ResumenAcademicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RolController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RolController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RolController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RolController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RolController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RubricaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RubricaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RubricaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RubricaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:RubricaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:RubricaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SesionOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SesionOrganizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SesionOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SesionOrganizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SesionOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SesionOrganizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SesionOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SesionOrganizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SesionOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SesionOrganizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SistemaCalendarioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SistemaCalendarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SistemaCalendarioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SistemaCalendarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SistemaCalendarioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SistemaCalendarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SistemaCalendarioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SistemaCalendarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SistemaCalendarioController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SistemaCalendarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudAdmisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudAdmisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudAdmisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudAdmisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudAdmisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudApoyoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudApoyoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudApoyoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudApoyoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudApoyoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudConsejeriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudConsejeriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudConsejeriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudConsejeriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SolicitudConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SolicitudConsejeriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoRecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoRecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoRecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoRecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:SoporteFisicoRecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoActividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoActividadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoActividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoActividadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoActividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoActividadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoActividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoActividadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoActividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoActividadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionInstitucionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAdmisionInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoApiAccesoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoApiAccesoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoApiAccesoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoApiAccesoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoApiAccesoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoApiAccesoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoApiAccesoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoApiAccesoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoApiAccesoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoApiAccesoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAutorController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAutorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAutorController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAutorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAutorController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAutorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAutorController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAutorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoAutorController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoAutorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEspacioAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEspacioAcademicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEspacioAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEspacioAcademicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEspacioAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEspacioAcademicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEspacioAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEspacioAcademicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEspacioAcademicoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEspacioAcademicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEstadoRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEstadoRolController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEstadoRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEstadoRolController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEstadoRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEstadoRolController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEstadoRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEstadoRolController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEstadoRolController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEstadoRolController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEventoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEventoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEventoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEventoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEventoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEventoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEventoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEventoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoEventoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoEventoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoIdentificacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoIdentificacionOrgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoIdentificacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoIdentificacionOrgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoIdentificacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoIdentificacionOrgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoIdentificacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoIdentificacionOrgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoIdentificacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoIdentificacionOrgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoInteractividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoInteractividadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoInteractividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoInteractividadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoInteractividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoInteractividadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoInteractividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoInteractividadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoInteractividadController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoInteractividadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoJornadaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoJornadaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoJornadaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoJornadaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoJornadaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoJornadaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoJornadaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoJornadaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoJornadaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoJornadaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoMatriculaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoMatriculaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoMatriculaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoMatriculaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoMatriculaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoMatriculaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoMatriculaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoMatriculaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoMatriculaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoMatriculaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionGrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionGrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionGrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionGrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionGrupoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionGrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionNotaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionNotaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionNotaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionNotaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionNotaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionNotaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionNotaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionNotaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionNotaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoObservacionNotaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoOrganizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoOrganizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoOrganizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoOrganizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoOrganizacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoOrganizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoParticipacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoParticipacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoParticipacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoParticipacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoParticipacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoParticipacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoParticipacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoParticipacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoParticipacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoParticipacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoProgramaLaboralController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoProgramaLaboralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoProgramaLaboralController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoProgramaLaboralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoProgramaLaboralController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoProgramaLaboralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoProgramaLaboralController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoProgramaLaboralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoProgramaLaboralController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoProgramaLaboralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoPruebaAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoPruebaAdmisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoPruebaAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoPruebaAdmisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoPruebaAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoPruebaAdmisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoPruebaAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoPruebaAdmisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoPruebaAdmisionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoPruebaAdmisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRecursoAprendizajeController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRecursoAprendizajeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRelacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRelacionOrgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRelacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRelacionOrgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRelacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRelacionOrgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRelacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRelacionOrgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRelacionOrgController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRelacionOrgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaCampusController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaCampusController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaCampusController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaCampusController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaCampusController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaCampusController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaCampusController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaCampusController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaCampusController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaCampusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoResidenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRiesgoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRiesgoRecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRiesgoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRiesgoRecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRiesgoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRiesgoRecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRiesgoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRiesgoRecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoRiesgoRecursoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoRiesgoRecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSeparacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSeparacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSeparacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSeparacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSeparacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSeparacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSeparacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSeparacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSeparacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSeparacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudApoyoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudApoyoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudApoyoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudApoyoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudApoyoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudApoyoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudConsejeriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudConsejeriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudConsejeriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudConsejeriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudConsejeriaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TipoSolicitudConsejeriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TituloOtorgadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TituloOtorgadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TituloOtorgadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TituloOtorgadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TituloOtorgadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TituloOtorgadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TituloOtorgadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TituloOtorgadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TituloOtorgadoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TituloOtorgadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TransferenciaListaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TransferenciaListaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TransferenciaListaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TransferenciaListaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TransferenciaListaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TransferenciaListaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TransferenciaListaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TransferenciaListaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:TransferenciaListaController"] = append(beego.GlobalControllerRouter["api_academica/controllers:TransferenciaListaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UnidadDedicacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UnidadDedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UnidadDedicacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UnidadDedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UnidadDedicacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UnidadDedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UnidadDedicacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UnidadDedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UnidadDedicacionController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UnidadDedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UnidadTiempoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UnidadTiempoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UnidadTiempoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UnidadTiempoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UnidadTiempoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UnidadTiempoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UnidadTiempoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UnidadTiempoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UnidadTiempoController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UnidadTiempoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UsuarioFinalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UsuarioFinalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UsuarioFinalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UsuarioFinalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UsuarioFinalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UsuarioFinalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UsuarioFinalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UsuarioFinalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_academica/controllers:UsuarioFinalController"] = append(beego.GlobalControllerRouter["api_academica/controllers:UsuarioFinalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
