export default function BooksList() {
  return (
    <body>
    <h1 className="title">Books List</h1>
    <table className="table is-striped is-narrow is-hoverable">
      <tr>
        <th>Nome</th>
        <th>ISBN</th>
        <th>Data</th>
        <th>Progresso</th>
      </tr>
      {/* {{range .}}
      <tr>
        <td>{{.Name}}</td>
        <td>{{.Isbn}}</td>
        <td>{{.Date}}</td>
        <td>{{.Progress}}</td>
      </tr>
      {{end}} */}
    </table>
  </body>
  );
}