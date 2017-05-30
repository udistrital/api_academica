package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ProgramaLaboral struct {
	Id                    int                    `orm:"column(id);pk"`
	FechaTitulacion       time.Time              `orm:"column(fecha_titulacion);type(date);null"`
	ParticipacionPrograma *ParticipacionPrograma `orm:"column(participacion_programa);rel(fk)"`
	TipoProgramaLaboral   *TipoProgramaLaboral   `orm:"column(tipo_programa_laboral);rel(fk)"`
	CategoriaTitulacion   *CategoriaTitulacion   `orm:"column(categoria_titulacion);rel(fk)"`
}

func (t *ProgramaLaboral) TableName() string {
	return "programa_laboral"
}

func init() {
	orm.RegisterModel(new(ProgramaLaboral))
}

// AddProgramaLaboral insert a new ProgramaLaboral into database and returns
// last inserted Id on success.
func AddProgramaLaboral(m *ProgramaLaboral) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProgramaLaboralById retrieves ProgramaLaboral by Id. Returns error if
// Id doesn't exist
func GetProgramaLaboralById(id int) (v *ProgramaLaboral, err error) {
	o := orm.NewOrm()
	v = &ProgramaLaboral{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProgramaLaboral retrieves all ProgramaLaboral matches certain condition. Returns empty list if
// no records exist
func GetAllProgramaLaboral(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProgramaLaboral))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ProgramaLaboral
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateProgramaLaboral updates ProgramaLaboral by Id and returns error if
// the record to be updated doesn't exist
func UpdateProgramaLaboralById(m *ProgramaLaboral) (err error) {
	o := orm.NewOrm()
	v := ProgramaLaboral{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProgramaLaboral deletes ProgramaLaboral by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProgramaLaboral(id int) (err error) {
	o := orm.NewOrm()
	v := ProgramaLaboral{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProgramaLaboral{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
