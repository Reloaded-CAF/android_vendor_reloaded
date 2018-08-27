package android

type Product_variables struct {
	Target_uses_color_metadata struct {
		Cppflags []string
	}
	Uses_generic_camera_parameter_library struct {
		Srcs []string
	}
	Uses_qcom_bsp_legacy struct {
		Cppflags []string
	}
	Uses_qti_camera_device struct {
		Cppflags    []string
		Shared_libs []string
	}
}

type ProductVariables struct {
	Specific_camera_parameter_library     *string `json:",omitempty"`
	Target_uses_color_metadata            *bool   `json:",omitempty"`
	Target_use_sdclang                    *bool   `json:",omitempty"`
	Uses_generic_camera_parameter_library *bool   `json:",omitempty"`
	Uses_qcom_bsp_legacy                  *bool   `json:",omitempty"`
	Uses_qti_camera_device                *bool   `json:",omitempty"`
}
