try {
    const init = (go, wasmFileLocation) => {
        return new Promise(
            (resolve, reject) => {
                console.log("Initializing Web Assembly...");
                if ("WebAssembly" in window) {
                    alert(JSON.stringify(window.WebAssembly));
                    window.WebAssembly.instantiateStreaming(fetch(wasmFileLocation), go.importObject).then(
                        (instanceDetails) => {
                            console.log("Success: Web Assembly initialized.");
                            console.log("Running Web Assembly...");
                            go.run(instanceDetails.instance).then(
                                () => {
                                    console.log("Success: Web Assembly ran.");
                                    resolve();
                                },
                                reject
                            );
                        },
                        reject
                    );
                } else {
                    reject("This browser does not support Web Assembly!");
                }
            }
        );
    };

    init(new Go(), "bin/out.wasm").catch(
        (error) => {
            console.error("Error: Web Assembly not initialized:");
            console.error(error);
        }
    );
} catch (e) {
    alert(e);
}
