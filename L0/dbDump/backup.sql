PGDMP  1                	    |            orders    16.3    16.3     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16509    orders    DATABASE     z   CREATE DATABASE orders WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';
    DROP DATABASE orders;
                postgres    false            �            1259    16510    orders    TABLE     �  CREATE TABLE public.orders (
    order_uid character varying NOT NULL,
    track_number character varying,
    entry character varying,
    delivery jsonb,
    payment jsonb,
    items jsonb,
    locale character varying,
    internal_signature character varying,
    customer_id character varying,
    delivery_service character varying,
    shardkey character varying,
    sm_id integer,
    date_created timestamp without time zone,
    oof_shard character varying
);
    DROP TABLE public.orders;
       public         heap    postgres    false            �          0    16510    orders 
   TABLE DATA           �   COPY public.orders (order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) FROM stdin;
    public          postgres    false    215                     2606    16516    orders orders_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_uid);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public            postgres    false    215            �   �  x�eRˎ�0]�_eS9��qW<h��h�ŀ����ZM�b;E��N�;E��%��{W�e�?*�%�u!��|�ȏw�>��7��׷���<�g}LWIʪ\�T�����aB�^�	B���6H�Y.�ÞV�q�y��۲�Rǽ5q���^?ĝ�ik��mb�uNy�����=t�Z;H�2��9%�r���v4a��:�{tN�6F����=�N9���#L��8�G��(��iP&4]�r.(�ǘ�ǹS�;T�᝵�o����hmxn$��=Nv��'妦�>��^��S�� �	Zs�U����m�k�-8��y��1�������9K���OZ��He�k�ׂe,nG�8Y�9�#n�Q$۽{�-D^��k7����Cc�A^���'K/��2�`2(<���]F�h��lUV��Xr^�? �?����/�ɜ     