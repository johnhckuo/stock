FROM ruby:2.6.6

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

RUN gem install bundler:2.1.4

COPY ./Gemfile /usr/src/app/Gemfile
COPY ./Gemfile.lock /usr/src/app/Gemfile.lock

RUN bundle install

COPY ./ /usr/src/app
