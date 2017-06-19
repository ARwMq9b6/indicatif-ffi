from setuptools import setup, find_packages

setup(
    name='indicatif',
    version='0.1.0',
    packages=find_packages(),
    include_package_data=True,
    zip_safe=False,
    platforms='any',
    setup_requires=["cffi>=1.0.0"],
    cffi_modules=["x/ffi_builder.py:ffibuilder"],
    install_requires=["cffi>=1.0.0"]
)
