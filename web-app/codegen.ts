import {CodegenConfig} from "@graphql-codegen/cli";

const config: CodegenConfig = {
    schema: "http://localhost:8080/graphql",
    documents: "src/types/*.tsx",
    ignoreNoDocuments: true,
    generates: {
        "./src/gql/": {
            preset: "client",
            plugins: []
        }
    }
}

export default config