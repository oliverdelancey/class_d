#!/usr/bin/env python3
"""setup"""

import setuptools


if __name__ == "__main__":
    with open("README.md", "r") as fh:
        LONG_DESCRIPTION = fh.read()

    setuptools.setup(
        name="get_license-oliversandli",
        version="0.0.1",
        author="Oliver Sandli",
        author_email="racerpingpong@gmail.com",
        description="Get a license from the GitHub API",
        long_description=LONG_DESCRIPTION,
        long_description_content_type="text/markdown",
        url="https://github.com/oliversandli/get_license",
        packages=["get_license"],
        scripts=["bin/get_license"],
        install_requires=[
            "inquirer",
            "requests"
        ],
        classifiers=[
            "Programming Language :: Python :: 3",
            "License :: OSI Approved :: MIT License",
            "Operating System :: Unix"
        ],
        python_requires='>=3.6'
    )
