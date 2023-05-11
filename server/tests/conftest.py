import json
import pytest

from urllib.parse import urljoin

@pytest.fixture(scope="package")
def env_config():
  with open("envs.json") as f:
    data = json.load(f)
  return data

@pytest.fixture(scope="package")
def env_routes(env_config):
  return env_config["routes"]

@pytest.fixture(scope="package")
def env_payloads(env_config):
  return env_config["data"]

@pytest.fixture(scope="package")
def base_url(env_config):
  return env_config["baseURL"]

@pytest.fixture(scope="package")
def login_route(env_config):
  return env_config["login_route"]

@pytest.fixture(scope="package")
def login_url(env_config):
  return env_config["login_url"]