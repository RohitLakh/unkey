import type { Env } from "@/pkg/env";
import { OpenAPIHono } from "@hono/zod-openapi";
import { prettyJSON } from "hono/pretty-json";
import { handleError, handleZodError } from "../errors";

export type HonoEnv = {
  Bindings: Env;
  Variables: {
    requestId: string;
  };
};
export function newApp() {
  const app = new OpenAPIHono<HonoEnv>({
    defaultHook: handleZodError,
  });

  app.onError(handleError);
  app.use(prettyJSON());

  app.doc("/openapi.json", {
    openapi: "3.0.0",
    info: {
      title: "Unkey Api",
      version: "1.0.0",
    },
    // @ts-expect-error - this is a bug in the types
    components: {
      securitySchemes: {
        BearerAuth: {
          type: "http",
          scheme: "bearer",
        },
      },
    },
    security: [{ BearerAuth: [] }],
    servers: [
      {
        url: "https://api.unkey.dev",
        description: "Production",
      },
    ],
  });

  // app.openAPIRegistry.registerComponent("securitySchemes", "BearerAuth", {
  //   type: "http",
  //   scheme: "bearer",
  // });
  return app;
}

export type App = ReturnType<typeof newApp>;
