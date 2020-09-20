import * as Table from "cli-table3";

export const endpointList = (baseUrl: string, routes: any): Table.Table => {
  const table = new Table({ head: ["method", "path"], style: { head: ["blue"] } });
  for (const key in routes) {
    if (routes.hasOwnProperty(key)) {
      const layer = routes[key];
      if (layer.route) {
        const route = layer.route;
        let _o: any = {};
        _o[route.stack[0].method] = [baseUrl + route.path];
        table.push(_o);
      }
    }
  }
  return table;
};
