package {{.Models}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
        log "github.com/sirupsen/logrus"
        {{range .Imports}}"{{.}}"{{end}}
)
{{end}}

{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}     {{Mapper $col.Name}}    {{Type $col}} {{Tag $table $col}}
{{end}}
}
{{end}}

{{range .Tables}}
func Insert{{Mapper .Name}}(m *{{Mapper .Name}}) error {
    var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "{{Mapper .Name}}": m}).Error("models.Insert{{Mapper .Name}} has an error")
		err = DBErr
	}
	return err
}

func Get{{Mapper .Name}}(id int) (m *{{Mapper .Name}}, err error) {
	var has bool
	m = &{{Mapper .Name}}{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.Get{{Mapper .Name}} has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func Update{{Mapper .Name}}(m *{{Mapper .Name}}) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "{{Mapper .Name}}": m}).Error("models.Update{{Mapper .Name}} has an error")
		err = DBErr
	}
	return err
}

func Delete{{Mapper .Name}}(id int) error {
    var err error
	if _, err := X.Id(id).Delete(new({{Mapper .Name}})); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.Delete{{Mapper .Name}} has an error")
		err = DBErr
	}
	return err
}
{{end}}
